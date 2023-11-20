-- users
-- name: get-user
-- Get a single user by username or email or phone
SELECT * FROM users WHERE
    CASE 
        WHEN $1 != '' THEN username = $1
        WHEN $2 != '' THEN email = $2
        WHEN $3 != '' THEN phone = $3
    END;

-- users
-- name: insert-user
-- Insert a user
INSERT INTO users(uuid, username, name, email, phone, pass_hash) 
VALUES($1, $2, $3, $4, $5, $6)
RETURNING uuid

-- expenses
-- name: insert-expense
INSERT INTO expenses(uuid, amount, title, split_mode, owner_uuid, group_uuid)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING uuid

-- splits
-- name: insert-split
INSERT INTO splits(uuid, expense_uuid, amount, owed_by_uuid, owed_to_uuid, settled, group_uuid)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING uuid