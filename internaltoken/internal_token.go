package internaltoken

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
)

const FilePath = "/data/rubix-service/data/internal_token.txt"

func CreateInternalToken() (string, error) {
	token := security.GenerateToken()
	_, err := file.WriteDataToFileAsString(FilePath, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetInternalToken() (string, error) {
	return file.ReadFile(FilePath)
}
