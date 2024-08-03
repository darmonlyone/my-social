BEGIN;

-- Drop triggers for the post table
DROP TRIGGER IF EXISTS post_updated_at ON post;

-- Drop triggers for the account table
DROP TRIGGER IF EXISTS account_updated_at ON account;

-- Drop function to update the updated_at field
DROP FUNCTION IF EXISTS set_current_time_updated_at;

-- Remove created_at and updated_at columns from the account table
ALTER TABLE account
DROP COLUMN created_at,
DROP COLUMN updated_at;

-- Remove created_at and updated_at columns from the post table
ALTER TABLE post
DROP COLUMN created_at,
DROP COLUMN updated_at;


END;