package externaltoken

import (
	"testing"
)

// Test are general test
func Test(t *testing.T) {
	// CreateToken
	externalToken, err := CreateExternalToken("token")
	if err != nil {
		t.Fatal("CreateExternalToken failed")
	}
	t.Log("CreateExternalToken success")

	// RegenerateExternalToken
	_, err = RegenerateExternalToken(externalToken.UUID)
	if err != nil {
		t.Fatal("RegenerateExternalToken failed")
	}
	t.Log("RegenerateExternalToken success")

	// RegenerateExternalToken
	_ = ValidateExternalToken(externalToken.Token)
	t.Log("ValidateExternalToken success")

	// BlockExternalToken
	_, err = BlockExternalToken(externalToken.UUID, true)
	if err != nil {
		t.Fatal("BlockExternalToken failed")
	}
	t.Log("BlockExternalToken success")

	// GetTokens
	_, err = GetExternalTokens()
	if err != nil {
		t.Fatal("GetExternalTokens failed")
	}
	t.Log("GetExternalTokens success")

	// DeleteToken
	_, err = DeleteExternalToken(externalToken.UUID)
	if err != nil {
		t.Fatal("DeleteExternalToken failed")
	}
	t.Log("DeleteExternalToken success")
}
