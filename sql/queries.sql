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
VALUES($1, $2, $3, $4, $5, $6);

-- name: insert-group
-- Inserts in groups
INSERT INTO groups(uuid, name)
VALUES($1, $2);

-- name: insert-user-group-mappings
INSERT INTO user_group_mappings(user_uuid, group_uuid)
VALUES($1, $2);

-- name: get-group
SELECT * FROM groups
INNER JOIN user_group_mappings ON (groups.uuid = user_group_mappings.group_uuid)
WHERE user_group_mappings.group_uuid = $1


-- name: get-groups-for-user
SELECT * FROM user_group_mappings
INNER JOIN groups ON (groups.uuid = user_group_mappings.group_uuid)
WHERE user_group_mappings.user_uuid = $1

-- name: insert-expense
INSERT INTO expenses(uuid, title, amount, owed_to_uuid, group_uuid)
VALUES ($1, $2, $3, $4, $5);

-- name: insert-user-expense-mapping
INSERT INTO user_expense_mappings(uuid, amount, owed_by_uuid, expense_uuid, group_uuid)
VALUES($1, $2, $3, $4, $5);

-- name: get-expense
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.uuid = user_expense_mappings.group_uuid)
WHERE user_expense_mappings.expense_uuid = $1

-- name: get-expenses-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.uuid = user_expense_mappings.group_uuid)
WHERE user_expense_mappings.group_uuid = $1

-- name: get-owed-by-expenses-for-user-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.uuid = user_expense_mappings.group_uuid)
WHERE user_expense_mappings.group_uuid = $1 AND user_expense_mappings.owed_by_uuid = $2;

-- name: get-owed-to-expenses-for-user-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.uuid = user_expense_mappings.group_uuid)
WHERE user_expense_mappings.group_uuid = $1 AND expenses.owed_to_uuid = $2;