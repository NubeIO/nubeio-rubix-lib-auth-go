package internaltoken

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/constants"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/security"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/utils/file"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

var internalToken *string

func GetInternalToken(withPrefix bool) string {
	if internalToken != nil {
		if withPrefix {
			return fmt.Sprintf("Internal %s", *internalToken)
		} else {
			return *internalToken
		}
	}
	filePath := file.GetDataFile(constants.InternalTokenFileName)
	f, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return ""
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Error(err)
		}
	}()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error(err)
	}
	it := string(bytes)
	internalToken = &it
	if withPrefix {
		return fmt.Sprintf("Internal %s", *internalToken)
	} else {
		return *internalToken
	}
}

func CreateInternalTokenIfDoesNotExist() {
	filePath := file.GetDataFile(constants.InternalTokenFileName)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		panic(err)
	}
	it, _ := file.ReadFile(filePath)
	if it != "" {
		return
	}
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Error(err)
	}
	token := security.GenerateToken()
	_, err = f.Write([]byte(token))
	if err != nil {
		log.Error(err)
	}
}
