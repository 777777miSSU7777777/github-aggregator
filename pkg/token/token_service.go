package token

import (
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"
)

// TOKEN_FILE name of file which stores personal access token.
const TOKEN_FILE = ".token"

// TokenService is service for working with access token.
type TokenService struct {
	apiToken     string
	tokenChecker tokenChecker
}

var tokenService *TokenService

func init() {
	tokenService = &TokenService{tokenChecker: tokenChecker{}}
}

// GetTokenService returns token service.
func GetTokenService() *TokenService {
	return tokenService
}

// SetTokenService sets token service.
func SetTokenService(ts *TokenService) {
	tokenService = ts
}

// TryLoadToken tries to load token from .token file in $HOME or current dir.
func (ts *TokenService) TryLoadToken() error {
	homeDir := os.Getenv("HOME")
	token, err := fileutil.ReadStringFromFile(homeDir + "/" + TOKEN_FILE)

	if token != "" {
		ts.SaveToken(token)
		return nil
	}

	token, err = fileutil.ReadStringFromFile(TOKEN_FILE)

	if token != "" {
		ts.SaveToken(token)
		return nil
	}

	ts.apiToken = ""

	if err != nil {
		return err
	}

	return &TokenNotFoundError{}
}

func (ts *TokenService) saveTokenFile() error {
	err := fileutil.WriteStringToFile(TOKEN_FILE, ts.apiToken)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TokenService) deleteTokenFile() error {
	err := os.Remove(TOKEN_FILE)

	if err != nil {
		return err
	}

	return nil
}

// SaveToken saves token to memory.
func (ts *TokenService) SaveToken(token string) error {
	valid, err := ts.tokenChecker.checkValidity(token)

	if err != nil {
		return err
	} else if valid {
		ts.apiToken = token

		ts.saveTokenFile()
	}

	return nil
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
