package internaltoken

import (
	"testing"
)

// Test are general test
func Test(t *testing.T) {
	// CreateInternalToken
	CreateInternalTokenIfDoesNotExist()
	t.Log("CreateInternalToken success")

	// GetInternalToken
	GetInternalToken(false)
	t.Log("GetInternalToken success")
}
