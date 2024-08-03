BEGIN;

-- Add created_at and updated_at columns to the account table
ALTER TABLE account
ADD COLUMN created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
ADD COLUMN updated_at TIMESTAMPTZ DEFAULT now() NOT NULL;

-- Add created_at and updated_at columns to the post table
ALTER TABLE post
ADD COLUMN created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
ADD COLUMN updated_at TIMESTAMPTZ DEFAULT now() NOT NULL;

-- Create a function to update the updated_at field
CREATE OR REPLACE FUNCTION set_current_time_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for the account table
CREATE TRIGGER account_updated_at BEFORE UPDATE ON account
FOR EACH ROW EXECUTE PROCEDURE set_current_time_updated_at();

-- Create triggers for the post table
CREATE TRIGGER post_updated_at BEFORE UPDATE ON post
FOR EACH ROW EXECUTE PROCEDURE set_current_time_updated_at();


END;