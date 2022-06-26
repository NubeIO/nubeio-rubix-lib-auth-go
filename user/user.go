package user

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	"regexp"
	"strings"
)

const FilePath = "/data/auth/user.txt"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func validateUsername(username string) bool {
	re, _ := regexp.Compile("^([A-Za-z0-9_-])+$")
	return re.FindString(username) != ""
}

func getUser() (*User, error) {
	data, err := file.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	user := strings.Split(data, ":")
	if len(user) < 2 {
		return nil, errors.New("user not found")
	}
	return &User{
		Username: user[0],
		Password: user[1],
	}, nil
}

func Login(user *User) (string, error) {
	q, err := getUser()
	if err != nil {
		return "", err
	}
	if q != nil && q.Username == user.Username && security.CheckPasswordHash(q.Password, user.Password) {
		return security.EncodeJwtToken(q.Username)
	}
	return "", errors.New("username and password combination is incorrect")
}

func GetUser() (*User, error) {
	q, err := getUser()
	if err != nil {
		return nil, err
	}
	if q != nil {
		q.Password = "******"
	}
	return q, nil
}

func CreateUser(user *User) (*User, error) {
	if !validateUsername(user.Username) {
		return nil, errors.New("username should be alphanumeric and can contain '_', '-'")
	}
	hashedPassword, err := security.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	_, err = file.WriteDataToFileAsString(FilePath, fmt.Sprintf("%s:%s", user.Username, hashedPassword))
	if err != nil {
		return nil, err
	}
	user.Password = "******"
	return user, nil
}
