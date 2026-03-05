-- 1. Удаляем триггер с таблицы
DROP TRIGGER IF EXISTS update_user_modtime ON users;

-- 2. Удаляем функцию обновления даты
DROP FUNCTION IF EXISTS update_modified_column();

-- 3. Удаляем таблицу
DROP TABLE IF EXISTS users;