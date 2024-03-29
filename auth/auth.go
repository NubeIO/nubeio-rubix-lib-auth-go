package auth

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/internaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/security"
	"net/http"
	"strings"
)

func GetAuthorization(request *http.Request) []string {
	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		return nil
	}
	return strings.SplitN(authorization, " ", 2)
}

func AuthorizeInternal(request *http.Request) bool {
	authorization := GetAuthorization(request)
	if len(authorization) == 2 && authorization[0] == "Internal" &&
		authorization[1] == internaltoken.GetInternalToken(false) {
		return true
	}
	return false
}

func AuthorizeExternal(request *http.Request) bool {
	authorization := GetAuthorization(request)
	if len(authorization) == 2 && authorization[0] == "External" {
		return externaltoken.ValidateExternalToken(authorization[1])
	}
	return false
}

func AuthorizeRoles(request *http.Request, roles ...string) (bool, *string, error) {
	authorization := GetAuthorization(request)
	if len(authorization) > 0 {
		authRole, err := GetAuthorizedRole(request)
		if err != nil {
			return false, nil, err
		}
		for _, role := range roles {
			if authRole == role {
				authorized, err := security.DecodeJwtToken(authorization[len(authorization)-1])
				return authorized, &authRole, err
			}
		}
	}
	return false, nil, errors.New("authorization header is missing")
}

func GetToken(request *http.Request) string {
	authorization := GetAuthorization(request)
	if len(authorization) > 0 {
		prefix := authorization[0]
		if len(authorization) == 2 && (prefix == "Internal" || prefix == "External") {
			return authorization[1]
		}
		return authorization[len(authorization)-1]
	}
	return ""
}

func GetAuthorizedUsername(request *http.Request) (string, error) {
	token := GetToken(request)
	return security.GetAuthorizedUsername(token)
}

func GetAuthorizedRole(request *http.Request) (string, error) {
	token := GetToken(request)
	return security.GetAuthorizedRole(token)
}
