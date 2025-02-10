package env

import (
	"errors"

	"github.com/joho/godotenv"
)

// file with settings for enviroment
const envLoc = ".env"

var env map[string]string

type AdminData struct {
	ID     int64
	GPTKey string
}

func Load() error {
	var err error
	if env, err = godotenv.Read(envLoc); err != nil {
		return err
	}
	return nil
}

// return token
func LoadTGToken() (string, error) {
	token, ok := env["TG_KEY"]
	if !ok {
		err := errors.New("telegram token not found in .env")
		return "", err
	}
	return token, nil
}

func LoadLocalPD() (string) {
	token:= env["LOCALHOST_PWD"]
	return token
}

func LoadLocalAI_Endpoint() (string) {
	token:= env["AI_ENDPOINT"]
	return token
}

func GetAdminToken() (string) {
	token := env["ADMIN_KEY"]
	return token
}
