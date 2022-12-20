-- name: CreateCollection :one
INSERT INTO collections (
  name, slug, notes, singleton, created_by
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateCollection :one
UPDATE collections SET name = ?, slug = ?, notes = ?, singleton = ?, schema = ?, listRule = ?, viewRule = ?, createRule = ?, updateRule = ?, deleteRule = ?, tags = ?, meta = ?, updated_at = ?, updated_by = ? 
WHERE id = ?
RETURNING *;

-- name: GetCollection :one
SELECT * FROM collections
WHERE id = ? LIMIT 1;

-- name: GetCollectionBySlug :one
SELECT * FROM collections
WHERE slug = ? LIMIT 1;

-- name: ListCollections :many
SELECT * FROM collections
ORDER BY name ASC;

-- name: DeleteCollection :execresult
DELETE FROM collections
WHERE id = ?;

