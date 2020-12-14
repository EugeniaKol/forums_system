package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "forums_sys",
		User:       "root",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "mysql://root:password@localhost/forums_sys?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
