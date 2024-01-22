-- name: CreateAccount :one
INSERT INTO accounts (owner, currency, balance)
VALUES ($1, $2, $3) RETURNING *;


-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id
    LIMIT $2
OFFSET $3;


-- name: UpdateAccounts :one
UPDATE accounts
SET balance = $1
WHERE id = $2 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;


