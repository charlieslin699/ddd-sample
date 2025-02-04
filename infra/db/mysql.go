package db

import (
	"time"

	gosqldrivermysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConn struct {
	db     *gorm.DB
	config gosqldrivermysql.Config
}

func NewMySQLConn(fns ...MySQLConnOptionFunc) (DBConn, error) {
	conn := &mysqlConn{
		config: gosqldrivermysql.Config{
			Timeout:      5 * time.Second, // 建立連線等待5秒
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 60 * time.Second,
			Net:          "tcp",
			ParseTime:    true,
		},
	}

	for _, fn := range fns {
		fn(conn)
	}

	var err error
	conn.db, err = gorm.Open(
		mysql.New(
			mysql.Config{
				DSNConfig: &conn.config,
			},
		),
	)

	return conn, err
}

func (conn *mysqlConn) DB() *gorm.DB {
	return conn.db
}

type MySQLConnOptionFunc func(conn *mysqlConn)

func WithUsername(username string) MySQLConnOptionFunc {
	return func(conn *mysqlConn) {
		conn.config.User = username
	}
}

func WithPassword(password string) MySQLConnOptionFunc {
	return func(conn *mysqlConn) {
		conn.config.Passwd = password
	}
}

func WithAddr(addr string) MySQLConnOptionFunc {
	return func(conn *mysqlConn) {
		conn.config.Addr = addr
	}
}

func WithDBName(dbName string) MySQLConnOptionFunc {
	return func(conn *mysqlConn) {
		conn.config.DBName = dbName
	}
}

func WithLocation(location time.Location) MySQLConnOptionFunc {
	return func(conn *mysqlConn) {
		conn.config.Loc = &location
	}
}
