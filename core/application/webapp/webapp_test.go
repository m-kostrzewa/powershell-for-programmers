package webapp_test

import (
	"net/http"
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/inmem"

	"github.com/stretchr/testify/assert"

	"github.com/m-kostrzewa/powershell-for-programmers/core/application/webapp"
)

func TestCanStartAndStop(t *testing.T) {
	questionsRepo := inmem.New()
	webapp := webapp.NewWebApp("../../..", questionsRepo)
	webapp.Serve(8080)
	webapp.Shutdown()
}

func TestServesCss(t *testing.T) {
	questionsRepo := inmem.New()
	webapp := webapp.NewWebApp("../../..", questionsRepo)
	webapp.Serve(8080)
	defer webapp.Shutdown()
	resp, err := http.Get("http://127.0.0.1:8080/static/style.css")
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
