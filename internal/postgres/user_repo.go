CREATE TABLE users (
    id  	 			TEXT 		PRIMARY KEY,
    user_name  	 		TEXT 		NOT NULL,
	user_role 	 		TEXT 		NOT NULL,
	joined_at 			TIMESTAMP	NOT NULL,
	changed_role_at		TIMESTAMP   
);

type PostgresUsersRepository struct{
	db     *sql.DB
}

func NewPostgresUserRepository(db *sql.DB)*PostgresUsersRepository{
	return &PostgresUsersRepository{
		db: db,
	}
}