package auth

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/internaltoken"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	"net/http"
	"strings"
)

func Authorize(request *http.Request) bool {
	authorization := strings.SplitN(request.Header.Get("Authorization"), " ", 2)
	if len(authorization) > 0 {
		// Internal Auth
		if len(authorization) == 2 && authorization[0] == "Internal" &&
			authorization[1] == internaltoken.GetInternalToken(false) {
			return true
		}
		// Token Auth
		if len(authorization) == 2 && authorization[0] == "External" {
			return externaltoken.ValidateExternalToken(authorization[1])
		}
		authorized, _ := security.DecodeJwtToken(authorization[len(authorization)-1])
		return authorized
	}
	return false
}
