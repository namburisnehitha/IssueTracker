CREATE TABLE comments(
    id          TEXT        PRIMARY KEY,
    issue_id    TEXT        REFERENCES issues(id), 
	user_id     TEXT        REFERENCES users(id),
	content     TEXT        NOT NULL,
	created_at  TIMESTAMP   NOT NULL,
	updated_at  TIMESTAMP   
);