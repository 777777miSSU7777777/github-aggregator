package encoding

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
