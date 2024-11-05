// models/user.go
package models

import (
    "database/sql"
    "errors"
)

type User struct {
    ID       int
    Username string
    Password string
}

func (u *User) Create(db *sql.DB) error {
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
    return err
}

func Authenticate(db *sql.DB, username, password string) (User, error) {
    var user User
    row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ? AND password = ?", username, password)
    err := row.Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return user, errors.New("invalid credentials")
    }
    return user, nil
}

