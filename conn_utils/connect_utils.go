package connect_utils

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/pbkdf2"
)

type DB_Information struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

var DB_info DB_Information

func (db_conn DB_Information) ConnectionString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		db_conn.User, db_conn.Password, db_conn.Server, db_conn.Port, db_conn.Database)
}

func (db_conn DB_Information) Open() *gorm.DB {
	db, err := gorm.Open("mssql", db_conn.ConnectionString())
	if err != nil {
		log.Fatal("Error opening Database: " + err.Error())
	}
	return db
}

func PwHash(password string, salt string) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 2048, 32, sha256.New)
}
