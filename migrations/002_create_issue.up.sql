CREATE TABLE issues (
    id  	 		    TEXT 		PRIMARY KEY,
    title 	 		    TEXT 		NOT NULL,
	issue_description	TEXT 		NOT NULL,
    issue_status        TEXT        NOT NULL,
	created_at 		    TIMESTAMP	NOT NULL,
	assignee_id         TEXT        REFERENCES users(id)
);