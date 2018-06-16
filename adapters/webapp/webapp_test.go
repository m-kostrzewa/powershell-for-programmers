package webapp_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
)

func TestCanStartAndStop(t *testing.T) {
	webapp := webapp.NewWebApp("../..", nil)
	webapp.Serve()
	webapp.Shutdown()
}

func TestServesCss(t *testing.T) {
	webapp := webapp.NewWebApp("../..", nil)
	webapp.Serve()
	defer webapp.Shutdown()
	resp, err := http.Get("http://127.0.0.1:8080/static/style.css")
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
