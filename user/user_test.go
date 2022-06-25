package user

import (
	"testing"
)

// Test are general test
func Test(t *testing.T) {
	// CreateUser
	_, err := CreateUser(&User{Username: "admin", Password: "admin"})
	if err != nil {
		t.Fatal("CreateUser failed")
	}
	t.Log("CreateUser success")

	// GetUser
	_, err = GetUser()
	if err != nil {
		t.Fatal("GetUser failed")
	}
	t.Log("GetUser success")

	// Login
	_, err = Login(&User{Username: "admin", Password: "admin"})
	if err != nil {
		t.Fatal("Login failed")
	}
	t.Log("Login success")
}
