package externaltoken

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	"strconv"
)

const FilePath = "/data/auth/external_token.csv"

type Token struct {
	Name    string `json:"name" `
	Token   string `json:"token" `
	Blocked bool   `json:"blocked" `
}

func mapToToken(records [][]string) []*Token {
	var tokens []*Token
	for _, value := range records {
		blocked, _ := strconv.ParseBool(value[2])
		token := Token{Name: value[0], Token: "******", Blocked: blocked}
		tokens = append(tokens, &token)
	}
	return tokens
}

func validateName(name string) error {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return err
	}
	for _, v := range records {
		if v[0] == name {
			return errors.New("duplicate name")
		}
	}
	return nil
}

func GetTokens() ([]*Token, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	return mapToToken(records), nil
}

func CreateToken(name string) (*Token, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	token := &Token{Name: name, Token: security.GenerateToken(), Blocked: false}
	records = append(records, []string{token.Name, token.Token, strconv.FormatBool(token.Blocked)})
	err = file.WriteCsvFile(FilePath, records)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func UpdateToken(name string, blocked bool) (*Token, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	for i, v := range records {
		if v[0] == name {
			v[2] = strconv.FormatBool(blocked)
			records[i] = v
			err = file.WriteCsvFile(FilePath, records)
			if err != nil {
				return nil, err
			}
			return &Token{Name: name, Token: "******", Blocked: blocked}, nil
		}
	}
	return nil, errors.New("token not found")

}

func DeleteToken(name string) error {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return err
	}
	for i, v := range records {
		if v[0] == name {
			records = append(records[:i], records[i+1:]...)
			return file.WriteCsvFile(FilePath, records)
		}
	}
	return errors.New("token not found")
}

func ValidateToken(token string) bool {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return false
	}
	for _, value := range records {
		blocked, _ := strconv.ParseBool(value[2])
		if value[1] == token && !blocked {
			return true
		}
	}
	return false
}
