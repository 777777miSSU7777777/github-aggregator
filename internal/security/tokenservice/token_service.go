package tokenservice

import (
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

var apiToken string

const TOKEN_FILE = ".token"

// TryLoadToken tries to load token from .token file in $HOME or current dir.
func TryLoadToken() {
	homeDir := os.Getenv("$HOME")
	token, err := fileutil.ReadStringFromFile(homeDir + "/" + TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}

	if token != "" {
		apiToken = token
		return
	}

	currentDir, err := os.Getwd()

	if err != nil {
		log.Warning.Println(err)
	}

	token, err = fileutil.ReadStringFromFile(currentDir + "/" + TOKEN_FILE)

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
	currentDir, err := os.Getwd()

	if err != nil {
		log.Warning.Println(err)
	}

	err = fileutil.WriteStringToFile(currentDir+"/"+TOKEN_FILE, apiToken)

	if err != nil {
		log.Warning.Println(err)
	}
}

func deleteTokenFile() {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Warning.Println(err)
	}

	err = os.Remove(currentDir + "/" + TOKEN_FILE)

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
