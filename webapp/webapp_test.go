package webapp_test

import (
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/webapp"
)

func TestCanStartAndStop(t *testing.T) {
	webapp := webapp.WebApp{}
	webapp.Serve("../templates")
	webapp.Shutdown()
}
