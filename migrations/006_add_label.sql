CREATE TABLE issue_labels (
    issue_id TEXT REFERENCES issues(id) ON DELETE CASCADE,
    label_id TEXT REFERENCES labels(id) ON DELETE CASCADE,
    PRIMARY KEY (issue_id, label_id)
);
