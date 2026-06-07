CREATE TABLE activities(
    id                       TEXT        PRIMARY KEY,
    issue_id                 TEXT        REFERENCES issues(id), 
	user_id                  TEXT        REFERENCES users(id),
	activity_description     TEXT        NOT NULL,
	created_at               TIMESTAMP   NOT NULL,
	activity_action          TEXT   
);