package externaltoken

import (
	"testing"
)

// Test are general test
func Test(t *testing.T) {
	// CreateToken
	_, err := CreateToken("token")
	if err != nil {
		t.Fatal("CreateToken failed")
	}
	t.Log("CreateToken success")

	// UpdateToken
	_, err = UpdateToken("token", true)
	if err != nil {
		t.Fatal("UpdateToken failed")
	}
	t.Log("UpdateToken success")

	// GetTokens
	_, err = GetTokens()
	if err != nil {
		t.Fatal("GetTokens failed")
	}
	t.Log("GetTokens success")

	// DeleteToken
	err = DeleteToken("token")
	if err != nil {
		t.Fatal("DeleteToken failed")
	}
	t.Log("DeleteToken success")
}
