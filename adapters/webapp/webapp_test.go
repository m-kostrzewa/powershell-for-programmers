package webapp_test

import (
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
)

func TestCanStartAndStop(t *testing.T) {
	webapp := webapp.WebApp{}
	webapp.Serve("../../templates")
	webapp.Shutdown()
}
