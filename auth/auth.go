package auth

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	"net/http"
	"strings"
)

func Authorize(request *http.Request) bool {
	auth := strings.SplitN(request.Header.Get("Authorization"), " ", 2)
	if len(auth) > 0 {
		// Internal Auth
		if len(auth) == 2 && auth[0] == "Internal" {
			return true
		}
		// Token Auth
		if len(auth) == 2 && auth[0] == "External" {
			return externaltoken.ValidateToken(auth[1])
		}
		authorized, err := security.DecodeJwtToken(auth[len(auth)-1])
		if err != nil {
			return false
		}
		return authorized
	}
	return false
}
