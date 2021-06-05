package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	name string
	age  int32
}

var (
	db *sql.DB
	e  error
)

func GetUser(n string) (user User, e error) {
	var name string
	var age int32

	row := db.QueryRow("SELECT name, age FROM users WHERE  name = ?", n)
	e = row.Scan(&name, &age)
	if e != nil {
		message := fmt.Sprintf("user %s not found", n)
		return User{}, errors.Wrap(sql.ErrNoRows, message)
	}

	user = User{name: name, age: age}
	return user, e
}

func init() {
	db, e = sql.Open("mysql",
		"{user}:{password}@tcp(127.0.0.1:3306)/{database}")
	if e != nil {
		panic(e)
	}
}
