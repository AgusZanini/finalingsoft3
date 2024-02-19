package config

import "os"

var (
	DBHOST = os.Getenv("DB_HOST")
	DBPORT = os.Getenv("DB_PORT")
	DBNAME = os.Getenv("DB_NAME")
	DBUSER = os.Getenv("DB_USER")
	DBPASS = os.Getenv("DB_PASS")
)
