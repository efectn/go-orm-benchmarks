-- name: CreateModel :exec
INSERT INTO models (NAME, title, fax, web, age, "right", counter)
VALUES ($1, $2, $3, $4, $5, $6, $7);
-- name: UpdateModel :exec
UPDATE models
SET name = $1,
    title = $2,
    fax = $3,
    web = $4,
    age = $5,
    "right" = $6,
    counter = $7
WHERE id = $8;
-- name: GetModel :one
SELECT *
FROM models
WHERE id = $1;
-- name: ListModels :many
SELECT *
FROM models
WHERE ID > 0
LIMIT 100;
-- name: InsertMulti :copyfrom
INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES
($1, $2, $3, $4, $5, $6, $7);
