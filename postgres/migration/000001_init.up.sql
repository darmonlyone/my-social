BEGIN;

-- Create Account table
CREATE TABLE IF NOT EXISTS account (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL
);

-- Create Post table
CREATE TABLE IF NOT EXISTS post (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_by UUID REFERENCES account(id) ON DELETE CASCADE NOT NULL,
    title TEXT NOT NULL,
    content TEXT
);

END;