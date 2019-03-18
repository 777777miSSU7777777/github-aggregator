package query

import (
	"fmt"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/bodyutil"
)


func QueryUser(tkn string)([]byte, error){
	resp, err := http.Get( fmt.Sprintf("%s%s?%s%s",constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn) )

	if err != nil {
		return nil , err
	}

	userBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return userBody, nil
}
