-- name: GetLinkByBlob :one
SELECT * FROM links WHERE short_link_blob = ?;

-- name: CreateShortLink :exec
INSERT INTO links (short_link_blob, long_link) 
VALUES (?, ?);
