package App

import (
	"github.com/Helladan/guakaTrack/Controllers"
	"io/ioutil"
	"encoding/json"
)

type secret struct {
	DatabaseHost     string `json:"database_host"`
	DatabaseName     string `json:"database_name"`
	DatabaseUsername string `json:"database_username"`
	DatabasePassword string `json:"database_password"`
}

// Get database logs
func getSecret() secret {
	raw, _ := ioutil.ReadFile("secret.json")

	var secret secret
	json.Unmarshal(raw, secret)

	return secret
}

var secretContent = getSecret()

var Database = Controllers.NewDatabase(
	secretContent.DatabaseHost,
	secretContent.DatabaseName,
	secretContent.DatabaseUsername,
	secretContent.DatabasePassword,
)
