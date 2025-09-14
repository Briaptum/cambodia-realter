-- Drop the trigger first
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop the function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_active;

-- Drop the users table
DROP TABLE IF EXISTS users;
