CREATE TABLE Comments (
    commentID SERIAL PRIMARY KEY,
    comment VARCHAR(256),
    publishedTime TIMESTAMP,
    parentCommentID SERIAL,
    urlString VARCHAR(128)
)