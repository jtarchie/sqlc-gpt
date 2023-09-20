-- name: Get :one
SELECT value
FROM keys
WHERE name = @name;
-- name: Substr :one
SELECT SUBSTR(
    value,
    IIF(@start < 0, @start, @start + 1),
    IIF(
      @end < 0,
      LENGTH(value) - @end,
      @start + @end + 1
    )
  )
FROM keys
WHERE name = @name;
-- name: ListLength :one
SELECT json_array_length(value) AS value
FROM keys
WHERE name = @name;
-- name: ListRange :many
SELECT json_each.value
FROM keys,
  json_each(keys.value)
WHERE keys.name = @name
  AND json_each.key >= IIF(@start >= 0, @start, json_array_length(keys.value) + @start)
  AND json_each.key <= IIF(@end >= 0, @end, json_array_length(keys.value) + @end);