package env

import (
	"errors"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

// file with settings for enviroment
const envLoc = ".env"

var env map[string]string

type AdminData struct {
	ID     int64
	GPTKey string
	localhost_password string
}

func Load() error {
	var err error
	if env, err = godotenv.Read(envLoc); err != nil {
		return err
	}
	return nil
}

// returns a map with admin data
func LoadAdminData() map[string]AdminData {
	adminData := make(map[string]AdminData)
	for admin, data := range env {
		switch admin {
		case "ADMIN_ID":

			id, err := strconv.ParseInt(data, 0, 64)
			if err != nil {
				log.Printf("admin id error parse: %s", data)
			}
			adminData["ADMIN_ID"] = AdminData{
				ID:     id,
				GPTKey: env["ADMIN_KEY"],
			}

		case "MINTY_ID":
			id, err := strconv.ParseInt(data, 0, 64)
			if err != nil {
				log.Printf("minty id error parse: %s", data)
			}
			adminData["MINTY_ID"] = AdminData{
				ID:     id,
				GPTKey: env["MINTY_KEY"],
			}

		case "OK_ID":
			id, err := strconv.ParseInt(data, 0, 64)
			if err != nil {
				log.Printf("ok id error parse: %s", data)
			}
			adminData["OK_ID"] = AdminData{
				ID:     id,
				GPTKey: env["OK_KEY"],
			}

		case "MURS_ID":
			id, err := strconv.ParseInt(data, 0, 64)
			if err != nil {
				log.Printf("murs id error parse: %s", data)
			}
			adminData["MURS_ID"] = AdminData{
				ID:     id,
				GPTKey: env["MURS_KEY"],
			}
		}
	}
	return adminData
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
