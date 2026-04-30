CREATE TABLE comments(
    id          TEXT        PRIMARY KEY,
    issue_id    TEXT        NOT NULL, 
	user_id     TEXT        NOT NULL,
	content     TEXT        NOT NULL,
	created_at  TIMESTAMP   NOT NULL,
	updated_at  TIMESTAMP   
);