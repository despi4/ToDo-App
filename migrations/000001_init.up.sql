CREATE TABLE users (
    id bigserial primary key,
    user_role varchar(50) not null,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email text unique not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create or replace function update_modified_column()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language 'plpgsql';

create trigger update_user_modtime
    before update on users
    for each row
    execute procedure update_modified_column();