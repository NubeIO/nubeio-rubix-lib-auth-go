package security

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/constants"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func handleSecretKey() (string, error) {
	filePath := file.GetDataFile(constants.SecretKeyFileName)
	secretKey, err := file.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	if secretKey == "" {
		secretKey := GenerateToken()
		_, err = file.WriteDataToFileAsString(filePath, secretKey)
		if err != nil {
			return "", err
		}
	}
	return secretKey, nil
}

func parseToken(token string) (*jwt.Token, error) {
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		secretKey, err := handleSecretKey()
		if err != nil {
			return nil, err
		}
		return []byte(secretKey), nil
	}
	return jwt.ParseWithClaims(token, jwt.MapClaims{}, key)
}

func getClaim(token string, claim string) (string, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		return "", err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return fmt.Sprintf("%v", claims[claim]), nil
	}
	return "", err
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

func EncodeJwtToken(userName string, role string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	atClaims["iat"] = time.Now().Unix()
	atClaims["sub"] = userName
	atClaims["role"] = role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secretKey, err := handleSecretKey()
	if err != nil {
		return "", err
	}
	return at.SignedString([]byte(secretKey))
}

func DecodeJwtToken(token string) (bool, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		return false, err
	}
	return parsedToken.Valid, err
}

func GetAuthorizedUsername(token string) (string, error) {
	return getClaim(token, "sub")
}

func GetAuthorizedRole(token string) (string, error) {
	return getClaim(token, "role")
}
