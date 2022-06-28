package externaltoken

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	"strconv"
)

const FilePath = "/data/auth/external_token.csv"

type ExternalToken struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name" `
	Token   string `json:"token" `
	Blocked bool   `json:"blocked" `
}

func makeUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%X%X", b[0:4], b[4:6])
	return
}

func mapToExternalToken(records [][]string) (externalTokens []*ExternalToken) {
	for _, record := range records {
		blocked, _ := strconv.ParseBool(record[3])
		externalTokens = append(externalTokens, &ExternalToken{UUID: record[0], Name: record[1], Token: "******", Blocked: blocked})
	}
	return externalTokens
}

func validateName(name string) error {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return err
	}
	for _, record := range records {
		if record[1] == name {
			return errors.New("name already exists")
		}
	}
	return nil
}

func GetExternalTokens() ([]*ExternalToken, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	return mapToExternalToken(records), nil
}

func CreateExternalToken(name string) (*ExternalToken, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	externalToken := &ExternalToken{UUID: makeUUID(), Name: name, Token: security.GenerateToken(), Blocked: false}
	records = append(records, []string{externalToken.UUID, externalToken.Name, externalToken.Token, strconv.FormatBool(externalToken.Blocked)})
	err = file.WriteCsvFile(FilePath, records)
	if err != nil {
		return nil, err
	}
	return externalToken, nil
}

func RegenerateExternalToken(uuid string) (*ExternalToken, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	for i, record := range records {
		if record[0] == uuid {
			record[2] = security.GenerateToken()
			records[i] = record
			err = file.WriteCsvFile(FilePath, records)
			if err != nil {
				return nil, err
			}
			blocked, _ := strconv.ParseBool(record[3])
			return &ExternalToken{UUID: uuid, Name: record[1], Token: record[2], Blocked: blocked}, nil
		}
	}
	return nil, errors.New("token not found")
}

func BlockExternalToken(uuid string, blocked bool) (*ExternalToken, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return nil, err
	}
	for i, record := range records {
		if record[0] == uuid {
			record[3] = strconv.FormatBool(blocked)
			records[i] = record
			err = file.WriteCsvFile(FilePath, records)
			if err != nil {
				return nil, err
			}
			return &ExternalToken{UUID: uuid, Name: record[1], Token: "******", Blocked: blocked}, nil
		}
	}
	return nil, errors.New("token not found")
}

func DeleteExternalToken(uuid string) (bool, error) {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return false, err
	}
	for i, record := range records {
		if record[0] == uuid {
			records = append(records[:i], records[i+1:]...)
			err := file.WriteCsvFile(FilePath, records)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, errors.New("token not found")
}

func ValidateExternalToken(token string) bool {
	records, err := file.ReadCsvFile(FilePath)
	if err != nil {
		return false
	}
	for _, record := range records {
		blocked, _ := strconv.ParseBool(record[3])
		if record[2] == token && !blocked {
			return true
		}
	}
	return false
}
