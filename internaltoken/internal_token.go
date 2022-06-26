package internaltoken

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/security"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

const FilePath = "/data/rubix-service/data/internal_token.txt"

func GetInternalToken(withPrefix bool) string {
	f, err := os.Open(FilePath)
	if err != nil {
		log.Error(err)
		return ""
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Error(err)
		}
	}()
	internalToken, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error(err)
	}
	if withPrefix {
		return fmt.Sprintf("Internal %s", string(internalToken))
	} else {
		return string(internalToken)
	}
}

func CreateInternalTokenIfDoesNotExist() {
	if err := os.MkdirAll(filepath.Dir(FilePath), 0755); err != nil {
		panic(err)
	}
	internalToken, _ := file.ReadFile(FilePath)
	if internalToken != "" {
		return
	}
	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Error(err)
	}
	token := security.GenerateToken()
	_, err = f.Write([]byte(token))
	if err != nil {
		log.Error(err)
	}
}
