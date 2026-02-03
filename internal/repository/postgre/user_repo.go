package postgre

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"todo-app/internal/domain/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Repository — это не “часть базы данных”
// Repository — это Adapter, который делает базу данных совместимой с бизнес-логикой

type UserRepo struct {
	db DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		db: *db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user users.User) (users.User, error) {
	sql := `
		INSERT INTO users (id, name, surname, email)
		values ($1, $2, $3, $4)
		returning id, name, surname, email, created_at, updated_at;
	`

	// INSERT INTO users..., добавляет новую строку в таблицу users и заполняет в ней только эти четыре колонки
	// created_at, updated_at они либони либо останутся пустыми (NULL), либо заполнятся значениями по умолчанию, если они настроены в схеме БД
	// values ($1, $2, $3, $4) (Prepared Statements) | Вместо того чтобы вставлять данные напрямую (например, 'Ivan'), ты оставляешь «дырки»
	//  Это единственный надежный способ защиты от SQL-инъекций
	// RETURNING id, name, surname, email, created_at, updated_at
	// говорит о том что какие данные я вставил в таблицу

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	var out users.User
	err := u.db.Pool.QueryRow(ctx, sql, user.ID, user.Name, user.Surname, user.Email).Scan(
		&out.ID,
		&out.Name,
		&out.Surname,
		&out.Email,
		&out.CreatedAt,
		&out.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return users.User{}, users.ErrEmailTaken
			}
		}
		return users.User{}, fmt.Errorf("repository CreateUser (email=%s): %w", user.Email, err)
	}

	return out, nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, ID uuid.UUID) (*users.User, error) {
	sql := `
		SELECT id, name, surname, email, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user users.User
	err := u.db.Pool.QueryRow(ctx, sql, ID).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, users.ErrNotFound
		}

		return nil, fmt.Errorf("repository GetUserByID (id=%s): %w", ID, err)
	}

	return &user, nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*users.User, error) {
	sql := `
		select id, name, surname, email, created_at, updated_at
		from users
		where email = $1;
	`

	var user users.User
	err := u.db.Pool.QueryRow(ctx, sql, email).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, users.ErrNotFound
		}

		return nil, fmt.Errorf("repository GetUserByEmail (email=%s): %w", email, err)
	}

	return &user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate users.UserUpdate) (users.User, error) {
	parts, args, pos := updateValidate(userUpdate)

	if len(parts) == 0 {
		user, err := u.GetUserByID(ctx, ID)
		if err != nil {
			return users.User{}, err
		}

		return *user, nil
	}

	parts = append(parts, "updated_at = now()")
	args = append(args, ID)

	wherePos := pos

	sql := fmt.Sprintf(`
		update users
		set %s
		where id = $%d
		returning id, name, surname, email, created_at, updated_at;
	`, strings.Join(parts, ", "), wherePos)

	var out users.User
	err := u.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&out.ID,
		&out.Name,
		&out.Surname,
		&out.Email,
		&out.CreatedAt,
		&out.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return users.User{}, users.ErrNotFound
		}

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return users.User{}, users.ErrEmailTaken
			}
		}

		return users.User{}, fmt.Errorf("repository UpdateUser (ID=%s): %w", ID, err)
	}

	return out, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, ID uuid.UUID) error {
	sql := `
		delete from users
		where id = $1;
	`

	cmd, err := u.db.Pool.Exec(ctx, sql, ID)
	if err != nil {
		return fmt.Errorf("repository DeleteUser (id=%s): %w", ID, err)
	}

	if cmd.RowsAffected() == 0 {
		return users.ErrNotFound
	}

	return nil
}

func updateValidate(userUpdate users.UserUpdate) (parts []string, args []any, position int) {
	parts = make([]string, 0, 4)
	args = make([]any, 0, 5)
	position = 1

	if userUpdate.Email != nil {
		parts = append(parts, fmt.Sprintf("email = $%d", position))
		args = append(args, *userUpdate.Email)
		position++
	}

	if userUpdate.Name != nil {
		parts = append(parts, fmt.Sprintf("name = $%d", position))
		args = append(args, *userUpdate.Name)
		position++
	}

	if userUpdate.Surname != nil {
		parts = append(parts, fmt.Sprintf("surname = $%d", position))
		args = append(args, *userUpdate.Surname)
		position++
	}

	return
}
