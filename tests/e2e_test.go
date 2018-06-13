package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/m-kostrzewa/powershell-for-programmers/webapp"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End2end Tests Suite")
}

var _ = Describe("Quiz", func() {
	var webapp webapp.WebApp
	BeforeEach(func() {
		webapp.Serve("../templates")
	})

	AfterEach(func() {
		webapp.Shutdown()
	})

	Context("/questions", func() {
		It("lists all available questions", func() {

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
