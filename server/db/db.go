package db

import (
	"database/sql"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

func (c *Connection) ConnectionURL() string {
	dbUrl := &url.URL{
		Scheme: "mysql",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbUrl.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
//	return sql.Open("mysql", c.ConnectionURL())
return sql.Open("mysql", "root:password@tcp(localhost:3306)/forums_sys")
}
