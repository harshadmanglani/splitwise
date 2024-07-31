-- users
DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
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
DROP INDEX IF EXISTS idx_users_username; CREATE UNIQUE INDEX idx_users_username ON users(username);

-- groups
DROP TABLE IF EXISTS groups CASCADE;
CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- user_group_mappings
DROP TABLE IF EXISTS user_group_mappings CASCADE;
CREATE TABLE user_group_mappings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (user_id),
    group_id INTEGER REFERENCES groups (group_id),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_group_id; CREATE UNIQUE INDEX idx_group_id ON user_group_mappings(group_id);
DROP INDEX IF EXISTS idx_user_id; CREATE UNIQUE INDEX idx_user_id ON user_group_mappings(user_id);

-- expenses
DROP TABLE IF EXISTS expenses CASCADE;
CREATE TABLE expenses (
    expense_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    owed_to_id INTEGER REFERENCES users (user_id),
    group_id INTEGER REFERENCES groups (group_id),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_owed_to; CREATE UNIQUE INDEX idx_owed_to ON expenses(owed_to_id);
DROP INDEX IF EXISTS idx_owed_to_per_group; CREATE UNIQUE INDEX idx_owed_to_per_group ON expenses(owed_to_id, group_id);

-- balances
DROP TABLE IF EXISTS balances CASCADE;
CREATE TABLE balances (
    balance_id SERIAL PRIMARY KEY,
    amount BIGINT NOT NULL,
    expense_id INTEGER REFERENCES expenses (expense_id),
    group_id INTEGER REFERENCES groups (group_id),
    settled BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_owed_by; CREATE UNIQUE INDEX idx_owed_by ON balances(owed_by_id);
DROP INDEX IF EXISTS idx_owed_by_per_group; CREATE UNIQUE INDEX idx_owed_by_per_group ON balances(owed_by_id, group_id);
DROP INDEX IF EXISTS idx_expense; CREATE UNIQUE INDEX idx_expense ON balances(expense_id);
