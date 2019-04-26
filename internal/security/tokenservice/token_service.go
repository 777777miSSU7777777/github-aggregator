package tokenservice

import (
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

var apiToken string

// TOKEN_FILE name of file which stores personal access token.
const TOKEN_FILE = ".token"

// TryLoadToken tries to load token from .token file in $HOME or current dir.
func TryLoadToken() {
	homeDir := os.Getenv("HOME")
	token, err := fileutil.ReadStringFromFile(homeDir + "/" + TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}

	if token != "" {
		apiToken = token
		return
	}

	token, err = fileutil.ReadStringFromFile(TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}

	if token != "" {
		apiToken = token
		return
	}

	apiToken = ""
}

func saveTokenFile() {
	err := fileutil.WriteStringToFile(TOKEN_FILE, apiToken)

	if err != nil {
		log.Warning.Println(err)
	}
}

func deleteTokenFile() {
	err := os.Remove(TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}
}

// SaveToken saves token to memory.
func SaveToken(token string) {
	apiToken = token

	saveTokenFile()
}

// GetToken return token.
func GetToken() string {
	return apiToken
}

// DeleteToken deletes token from memory.
func DeleteToken() {
	apiToken = ""

	deleteTokenFile()
}
