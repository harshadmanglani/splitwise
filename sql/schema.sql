-- users
DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    name TEXT NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_users_email; CREATE UNIQUE INDEX idx_users_email ON users(LOWER(email));
DROP INDEX IF EXISTS idx_users_phone; CREATE UNIQUE INDEX idx_users_phone ON users(phone);