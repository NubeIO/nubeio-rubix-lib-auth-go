package security

import (
	"crypto/rand"
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SecretKeyFilePath = "/data/auth/secret.txt"

func handleSecretKey() (string, error) {
	secretKey, err := file.ReadFile(SecretKeyFilePath)
	if err != nil {
		return "", err
	}
	if secretKey == "" {
		secretKey := GenerateToken()
		_, err = file.WriteDataToFileAsString(SecretKeyFilePath, secretKey)
		if err != nil {
			return "", err
		}
	}
	return secretKey, nil
}

func GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GenerateToken() string {
	uniqueKey := make([]byte, 16)
	_, _ = rand.Read(uniqueKey)
	hashedPassword, err := bcrypt.GenerateFromPassword(uniqueKey, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func EncodeJwtToken(userName string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	atClaims["iat"] = time.Now().Unix()
	atClaims["sub"] = userName
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secretKey, err := handleSecretKey()
	if err != nil {
		return "", err
	}
	return at.SignedString([]byte(secretKey))
}

func DecodeJwtToken(token string) (bool, error) {
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		secretKey, err := handleSecretKey()
		if err != nil {
			return false, err
		}
		return []byte(secretKey), nil
	}
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, key)
	if err != nil {
		return false, err
	}
	return parsedToken.Valid, err
}
