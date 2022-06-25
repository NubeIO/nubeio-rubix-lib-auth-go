package internaltoken

import (
	"testing"
)

// Test are general test
func Test(t *testing.T) {
	// CreateInternalToken
	_, err := CreateInternalToken()
	if err != nil {
		t.Fatal("CreateInternalToken failed")
	}
	t.Log("CreateInternalToken success")

	// GetInternalToken
	_, err = GetInternalToken()
	if err != nil {
		t.Fatal("GetInternalToken failed")
	}
	t.Log("GetInternalToken success")
}
