package tokenservice

// import (
// 	"os"
// 	"testing"

// 	"github.com/777777miSSU7777777/github-aggregator/pkg/io/fileutil"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestTryLoadToken__HomeDir__NotEmpty(t *testing.T) {
// 	fileutil.WriteStringToFile(os.Getenv("HOME")+"/"+TOKEN_FILE, "123")

// 	GetTokenService().TryLoadToken()

// 	assert.NotEmpty(t, GetTokenService().apiToken)

// 	os.Remove(os.Getenv("HOME") + "/" + TOKEN_FILE)

// 	GetTokenService().apiToken = ""
// }

// func TestTryLoadToken__CurrentDir__NotEmpty(t *testing.T) {
// 	fileutil.WriteStringToFile(TOKEN_FILE, "123")

// 	GetTokenService().TryLoadToken()

// 	assert.NotEmpty(t, GetTokenService().apiToken)

// 	os.Remove(TOKEN_FILE)

// 	GetTokenService().apiToken = ""
// }

// func TestTryLoadToken__NoTokenFiles__Empty(t *testing.T) {
// 	GetTokenService().TryLoadToken()

// 	assert.Empty(t, GetTokenService().apiToken)
// }

// func TestSaveTokenFile__SameString__Equals(t *testing.T) {
// 	GetTokenService().apiToken = "123"

// 	GetTokenService().saveTokenFile()

// 	token, _ := fileutil.ReadStringFromFile(TOKEN_FILE)

// 	assert.Equal(t, GetTokenService().apiToken, token)

// 	GetTokenService().apiToken = ""
// }

// func TestSaveTokenFile__DifferentStrings__NotEquals(t *testing.T) {
// 	GetTokenService().apiToken = "123"

// 	GetTokenService().saveTokenFile()

// 	GetTokenService().apiToken = "321"

// 	token, _ := fileutil.ReadStringFromFile(TOKEN_FILE)

// 	assert.NotEqual(t, GetTokenService().apiToken, token)

// 	GetTokenService().apiToken = ""
// }

// func TestDeleteTokenFile__TokenFileDeleted__Error(t *testing.T) {
// 	GetTokenService().saveTokenFile()

// 	GetTokenService().deleteTokenFile()

// 	_, err := os.Open(TOKEN_FILE)

// 	assert.Error(t, err)
// }

// func TestSaveToken__SameString__Equals(t *testing.T) {
// 	GetTokenService().SaveToken("123")

// 	assert.Equal(t, "123", GetTokenService().apiToken)

// 	GetTokenService().apiToken = ""
// }

// func TestSaveToken__DifferentStrings__NotEquals(t *testing.T) {
// 	GetTokenService().SaveToken("123")

// 	assert.NotEqual(t, "321", GetTokenService().apiToken)
// }

// func TestGetToken__SameString__Equals(t *testing.T) {
// 	GetTokenService().apiToken = "123"

// 	assert.Equal(t, GetTokenService().apiToken, GetTokenService().GetToken())
// }

// func TestGetToken__DifferentStrings__NotEquals(t *testing.T) {
// 	GetTokenService().apiToken = "123"

// 	assert.NotEqual(t, "321", GetTokenService().GetToken())
// }

// func TestDeleteToken__DeletedToken__Empty(t *testing.T) {
// 	GetTokenService().apiToken = "123"

// 	GetTokenService().DeleteToken()

// 	assert.Empty(t, GetTokenService().apiToken)
// }
