-- users
DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
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

-- groups
DROP TABLE IF EXISTS groups;
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    title VARCHAR(255) NOT NULL,
    simplify boolean DEFAULT false,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- expenses
DROP TYPE split_mode IF EXISTS CASCADE; CREATE TYPE split_mode AS ENUM ('EQUAL', 'PERCENTAGE', 'AMOUNT');
DROP TABLE IF EXISTS expenses CASCADE;
CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    amount bigint NOT NULL CHECK (amount > 0),
    title VARCHAR(255) NOT NULL,
    split_mode split_mode DEFAULT 'EQUAL',
    owner_uuid uuid NOT NULL REFERENCES users(uuid),
    group_uuid uuid REFERENCES groups(uuid),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- splits
DROP TABLE IF EXISTS splits CASCADE;
CREATE TABLE splits (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    expense_uuid uuid NOT NULL REFERENCES expenses(uuid),
    group_uuid uuid REFERENCES groups(uuid),
    amount bigint NOT NULL CHECK (amount > 0),
    owed_by_uuid uuid NOT NULL REFERENCES users(uuid),
    owed_to_uuid uuid NOT NULL REFERENCES users(uuid),
    settled boolean DEFAULT false,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
