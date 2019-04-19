package tokenservice

import (
	"os"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"

	"github.com/stretchr/testify/assert"
)

func TestTryLoadToken__HomeDir__NotEmpty(t *testing.T) {
	fileutil.WriteStringToFile(os.Getenv("HOME")+"/"+TOKEN_FILE, "123")

	TryLoadToken()

	assert.NotEmpty(t, apiToken)

	os.Remove(os.Getenv("HOME") + "/" + TOKEN_FILE)

	apiToken = ""
}

func TestTryLoadToken__CurrentDir__NotEmpty(t *testing.T) {
	fileutil.WriteStringToFile(TOKEN_FILE, "123")

	TryLoadToken()

	assert.NotEmpty(t, apiToken)

	os.Remove(TOKEN_FILE)

	apiToken = ""
}

func TestTryLoadToken__NoTokenFiles__Empty(t *testing.T) {
	TryLoadToken()

	assert.Empty(t, apiToken)
}

func TestSaveTokenFile__SameString__Equals(t *testing.T) {
	apiToken = "123"

	saveTokenFile()

	token, _ := fileutil.ReadStringFromFile(TOKEN_FILE)

	assert.Equal(t, apiToken, token)

	apiToken = ""
}

func TestSaveTokenFile__DifferentStrings__NotEquals(t *testing.T) {
	apiToken = "123"

	saveTokenFile()

	apiToken = "321"

	token, _ := fileutil.ReadStringFromFile(TOKEN_FILE)

	assert.NotEqual(t, apiToken, token)

	apiToken = ""
}

func TestDeleteTokenFile__TokenFileDeleted__Error(t *testing.T) {
	saveTokenFile()

	deleteTokenFile()

	_, err := os.Open(TOKEN_FILE)

	assert.Error(t, err)
}

func TestSaveToken__SameString__Equals(t *testing.T) {
	SaveToken("123")

	assert.Equal(t, "123", apiToken)

	apiToken = ""
}

func TestSaveToken__DifferentStrings__NotEquals(t *testing.T) {
	SaveToken("123")

	assert.NotEqual(t, "321", apiToken)
}

func TestGetToken__SameString__Equals(t *testing.T) {
	apiToken = "123"

	assert.Equal(t, apiToken, GetToken())
}

func TestGetToken__DifferentStrings__NotEquals(t *testing.T) {
	apiToken = "123"

	assert.NotEqual(t, "321", GetToken())
}

func TestDeleteToken__DeletedToken__Empty(t *testing.T) {
	apiToken = "123"

	DeleteToken()

	assert.Empty(t, apiToken)
}
