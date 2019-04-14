package login

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestLoginRender(t *testing.T) {
	view.SetTemplates(template.Must(template.ParseGlob("../../../web/templates/login.gohtml")))

	rw := httptest.NewRecorder()

	req := &http.Request{}

	Render(rw, req)

	assert.NotEmpty(t, rw.Body)
}
