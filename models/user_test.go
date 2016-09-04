package models_test

import (
	"testing"

	. "./"
)

func TestAuthenticateFailsWithIncorrectPasswor(t *testing.T) {
	user := NewUser("gopher@golang.go", "password")
	if user.Authenticate("wrong") {
		t.Errorf("Authenticate should return false for incorrect password")
	}
}

func TestAuthenticateSucceedsWithCorrectPasswor(t *testing.T) {
	user := NewUser("gopher@golang.go", "password")
	if !user.Authenticate("password") {
		t.Errorf("Authenticate should return true for correct password")
	}
}
