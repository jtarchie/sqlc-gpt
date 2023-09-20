-- name: DeleteMany :batchone
DELETE FROM keys
WHERE name IN (@names)
RETURNING value;
-- name: GetMany :batchmany
SELECT name,
  value
FROM keys
WHERE name IN (@names);