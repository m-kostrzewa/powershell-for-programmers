package e2e_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End2end Tests Suite")
}

var _ = Describe("Quiz", func() {
	var ts *httptest.Server
	BeforeEach(func() {
		app := webapp.NewWebApp("../templates")
		ts = httptest.NewServer(app.Mux)
	})

	AfterEach(func() {
		ts.Close()
	})

	Context("/questions", func() {
		It("lists all available questions", func() {
			resp, err := http.Get(ts.URL + "/questions")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})

	Context("/questions/0", func() {
		It("shows the question", func() {

		})
	})

	Context("/questions/answer", func() {
		It("shows congrats if the selected answer is correct", func() {

		})

		It("shows condolences if the selected answer is incorrect", func() {

		})
	})
})
