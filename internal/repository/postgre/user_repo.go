package postgre

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"todo-app/internal/db"
	userdomain "todo-app/internal/domain/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Repository — это не “часть базы данных”
// Repository — это Adapter, который делает базу данных совместимой с бизнес-логикой

type UserRepo struct {
	db db.DB
}

func NewUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{
		db: *db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user userdomain.User) (userdomain.User, error) {
	query := `
		INSERT INTO users (id, first_name, last_name, email, user_role, password_hash)
		values ($1, $2, $3, $4, $5, $6)
		returning id, first_name, last_name, email, created_at, updated_at, role, password_hash;
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

	if user.Role == "" {
		user.Role = userdomain.RoleUser
	}

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	var out userdomain.User
	err := u.db.Pool.QueryRow(ctx, query, user.ID, user.Name, user.Surname, user.Email, user.Role, user.PasswordHash).Scan(
		&out.ID,
		&out.Name,
		&out.Surname,
		&out.Email,
		&out.CreatedAt,
		&out.UpdatedAt,
		&out.Role,
		&out.PasswordHash,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return userdomain.User{}, userdomain.ErrEmailTaken
			}
		}
		return userdomain.User{}, fmt.Errorf("repository CreateUser (email=%s): %w", user.Email, err)
	}

	return out, nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, ID uuid.UUID) (*userdomain.User, error) {
	query := `
		SELECT id, first_name, last_name, email, created_at, updated_at, role
		FROM users
		WHERE id = $1;
	`

	var user userdomain.User
	err := u.db.Pool.QueryRow(ctx, query, ID).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userdomain.ErrNotFound
		}

		return nil, fmt.Errorf("repository GetUserByID (id=%s): %w", ID, err)
	}

	return &user, nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	query := `
		select id, first_name, last_name, email, created_at, updated_at, role
		from users
		where email = $1;
	`

	email = strings.ToLower(strings.TrimSpace(email))

	var user userdomain.User
	err := u.db.Pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userdomain.ErrNotFound
		}

		return nil, fmt.Errorf("repository GetUserByEmail (email=%s): %w", email, err)
	}

	return &user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate userdomain.UpdateUser) (userdomain.User, error) {
	parts, args, pos := updateValidate(userUpdate)

	if len(parts) == 0 {
		user, err := u.GetUserByID(ctx, ID)
		if err != nil {
			return userdomain.User{}, err
		}

		return *user, nil
	}

	parts = append(parts, "updated_at = now()")
	args = append(args, ID)

	wherePos := pos

	query := fmt.Sprintf(`
		update users
		set %s
		where id = $%d
		returning id, first_name, last_name, email, created_at, updated_at, role;
	`, strings.Join(parts, ", "), wherePos)

	var out userdomain.User
	err := u.db.Pool.QueryRow(ctx, query, args...).Scan(
		&out.ID,
		&out.Name,
		&out.Surname,
		&out.Email,
		&out.CreatedAt,
		&out.UpdatedAt,
		&out.Role,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return userdomain.User{}, userdomain.ErrNotFound
		}

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return userdomain.User{}, userdomain.ErrEmailTaken
			}
		}

		return userdomain.User{}, fmt.Errorf("repository UpdateUser (ID=%s): %w", ID, err)
	}

	return out, nil
}

func updateValidate(userUpdate userdomain.UpdateUser) (parts []string, args []any, position int) {
	parts = make([]string, 0, 4)
	args = make([]any, 0, 5)
	position = 1

	if userUpdate.Email != nil {
		parts = append(parts, fmt.Sprintf("email = $%d", position))
		args = append(args, *userUpdate.Email)
		position++
	}

	if userUpdate.Name != nil {
		parts = append(parts, fmt.Sprintf("first_name = $%d", position))
		args = append(args, *userUpdate.Name)
		position++
	}

	if userUpdate.Surname != nil {
		parts = append(parts, fmt.Sprintf("last_name = $%d", position))
		args = append(args, *userUpdate.Surname)
		position++
	}

	return
}

func (u *UserRepo) UpdatePasswordHash(ctx context.Context, ID uuid.UUID, new_hash userdomain.PasswordHash) error {
	query := `
		update users
		set password_hash = $1, updated_at = now()
		where id = $2
	`

	cmdTag, err := u.db.Pool.Exec(ctx, query, new_hash, ID)
	if err != nil {
		return fmt.Errorf("repository UpdateUserPassword (id=%s): %w", ID, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return userdomain.ErrNotFound
	}

	return nil
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
		return userdomain.ErrNotFound
	}

	return nil
}
