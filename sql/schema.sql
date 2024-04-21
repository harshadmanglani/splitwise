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
DROP TABLE IF EXISTS groups CASCADE;
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    name VARCHAR(255) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- user_group_mappings
DROP TABLE IF EXISTS user_group_mappings CASCADE;
CREATE TABLE user_group_mappings (
    id SERIAL PRIMARY KEY,
    user_uuid uuid NOT NULL,
    group_uuid uuid NOT NULL,

    FOREIGN KEY (user_uuid) REFERENCES users (uuid),
    FOREIGN KEY (group_uuid) REFERENCES groups (uuid),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_group_uuid; CREATE UNIQUE INDEX idx_group_uuid ON user_group_mappings(group_uuid);
DROP INDEX IF EXISTS idx_user_uuid; CREATE UNIQUE INDEX idx_user_uuid ON user_group_mappings(user_uuid);

-- expenses
DROP TABLE IF EXISTS expenses CASCADE;
CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    title VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    owed_to_uuid uuid NOT NULL,
    group_uuid uuid,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    FOREIGN KEY (owed_to_uuid) REFERENCES users (uuid),
    FOREIGN KEY (group_uuid) REFERENCES groups (uuid),
);
DROP INDEX IF EXISTS idx_owed_to; CREATE UNIQUE INDEX idx_owed_to ON expenses(owed_to_uuid);
DROP INDEX IF EXISTS idx_owed_to_per_group; CREATE UNIQUE INDEX idx_owed_to_per_group ON expenses(owed_to_uuid, group_uuid);

-- balances
DROP TABLE IF EXISTS balances CASCADE;
CREATE TABLE balances (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    amount BIGINT NOT NULL,
    owed_by_uuid uuid NOT NULL,
    expense_uuid uuid NOT NULL,
    group_uuid uuid,
    settled BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    FOREIGN KEY (owed_by_uuid) REFERENCES users (uuid),
    FOREIGN KEY (group_uuid) REFERENCES groups (uuid),
    FOREIGN KEY (expense_uuid) REFERENCES expenses (uuid),
);
DROP INDEX IF EXISTS idx_owed_by; CREATE UNIQUE INDEX idx_owed_by ON balances(owed_by_uuid);
DROP INDEX IF EXISTS idx_owed_by_per_group; CREATE UNIQUE INDEX idx_owed_by_per_group ON balances(owed_by_uuid, group_uuid);
DROP INDEX IF EXISTS idx_expense; CREATE UNIQUE INDEX idx_expense ON balances(expense_uuid);
