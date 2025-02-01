-- users
-- name: get-user
-- Get a single user by user_id
SELECT * FROM users WHERE
    CASE 
        WHEN $1 != '' THEN username = $1
        WHEN $2 != '' THEN email = $2
        WHEN $3 != '' THEN phone = $3
        WHEN $4 != '' THEN user_id = $4
    END;

-- users
-- name: insert-user
-- Insert a user
INSERT INTO users(user_id, username, name, email, phone, pass_hash) 
VALUES($1, $2, $3, $4, $5, $6);

-- name: insert-group
-- Inserts in groups
INSERT INTO groups(group_id, name)
VALUES($1, $2);

-- name: insert-user-group-mappings
INSERT INTO user_group_mappings(user_id, group_id)
VALUES($1, $2);

-- name: get-group
SELECT * FROM groups
INNER JOIN user_group_mappings ON (groups.group_id = user_group_mappings.group_id)
WHERE user_group_mappings.group_id = $1


-- name: get-groups-for-user
SELECT * FROM user_group_mappings
INNER JOIN groups ON (groups.id = user_group_mappings.group_id)
WHERE user_group_mappings.user_id = $1

-- name: insert-expense
INSERT INTO expenses(id, title, amount, owed_to_id, group_id)
VALUES ($1, $2, $3, $4, $5);

-- name: insert-user-expense-mapping
INSERT INTO user_expense_mappings(id, amount, owed_by_id, expense_id, group_id)
VALUES($1, $2, $3, $4, $5);

-- name: get-expense
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.expense_id = user_expense_mappings.group_id)
WHERE user_expense_mappings.expense_id = $1

-- name: get-expenses-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.expense_id = user_expense_mappings.group_id)
WHERE user_expense_mappings.group_id = $1

-- name: get-owed-by-expenses-for-user-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.expense_id = user_expense_mappings.group_id)
WHERE user_expense_mappings.group_id = $1 AND user_expense_mappings.owed_by_id = $2;

-- name: get-owed-to-expenses-for-user-in-group
SELECT * FROM expenses
INNER JOIN user_expense_mappings ON (expenses.id = user_expense_mappings.group_id)
WHERE user_expense_mappings.group_id = $1 AND expenses.owed_to_id = $2;