-- name: CreateComment :exec
INSERT INTO Comments (
    comment, publishedTime, urlString
)   VALUES (
    $1, $2, $3
);

-- name: CreateReplyComment :exec
INSERT INTO Comments (
    comment, publishedTime, parentCommentID, urlString
) VALUES (
    $1, $2, $3, $4
);

-- name: RetrieveComments :many
SELECT * FROM Comments
WHERE parentCommentID IS NULL;

-- name: RetrieveReplies :many
SELECT * FROM Comments
WHERE parentCommentID = $1;