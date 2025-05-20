package database

import "database/sql"

func ConnectDB() {
	postgresInfo =
		sql.Open("postgres")
}
