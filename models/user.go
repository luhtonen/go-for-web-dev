package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	Username string `db:"username"`
	Secret   []byte `db:"secret"`
}

// NewUser creates new database user
func NewUser(un, pw string) *User {
	secret, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return &User{
		Username: un,
		Secret:   secret,
	}
}

// Authenticate verify user's password
func (u *User) Authenticate(pw string) bool {
	return bcrypt.CompareHashAndPassword(u.Secret, []byte(pw)) == nil
}
