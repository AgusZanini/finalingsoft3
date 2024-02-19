package config

import (
	"os"
)

var (
	DBHOST = os.Getenv("INSTANCE_CONNECTION_NAME")
	DBPORT = os.Getenv("DB_PORT")
	DBNAME = os.Getenv("DB_NAME")
	DBUSER = os.Getenv("DB_USER")
	DBPASS = os.Getenv("DB_PASS")
)

/*
func DefinePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
*/
