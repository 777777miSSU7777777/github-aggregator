package tokenservice

import (
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

// TOKEN_FILE name of file which stores personal access token.
const TOKEN_FILE = ".token"

type TokenService struct {
	apiToken     string
	tokenChecker tokenChecker
}

var tokenService *TokenService

func init() {
	tokenService = &TokenService{tokenChecker: tokenChecker{}}
}

func GetTokenService() *TokenService {
	return tokenService
}

// TryLoadToken tries to load token from .token file in $HOME or current dir.
func (ts *TokenService) TryLoadToken() {
	homeDir := os.Getenv("HOME")
	token, err := fileutil.ReadStringFromFile(homeDir + "/" + TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}

	if token != "" {
		ts.SaveToken(token)
		return
	}

	token, err = fileutil.ReadStringFromFile(TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}

	if token != "" {
		ts.SaveToken(token)
		return
	}

	ts.apiToken = ""
}

func (ts *TokenService) saveTokenFile() {
	err := fileutil.WriteStringToFile(TOKEN_FILE, ts.apiToken)

	if err != nil {
		log.Warning.Println(err)
	}
}

func (ts *TokenService) deleteTokenFile() {
	err := os.Remove(TOKEN_FILE)

	if err != nil {
		log.Warning.Println(err)
	}
}

// SaveToken saves token to memory.
func (ts *TokenService) SaveToken(token string) {
	valid, err := ts.tokenChecker.checkValidity(token)

	if err != nil {
		log.Warning.Println(err)
		return
	} else if valid {
		ts.apiToken = token

		ts.saveTokenFile()
	}
}

// GetToken return token.
func (ts TokenService) GetToken() string {
	return ts.apiToken
}

// DeleteToken deletes token from memory.
func (ts *TokenService) DeleteToken() {
	ts.apiToken = ""

	ts.deleteTokenFile()
}
