package e2e_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
	"github.com/m-kostrzewa/powershell-for-programmers/core/question"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End2end Tests Suite")
}

var _ = Describe("Quiz", func() {
	var ts *httptest.Server

	questions := []question.Question{
		*question.New(question.NextQuestionID(),
			"Lexical scope",
			"Does Powershell do X?",
			"Some pseudocode here....",
			[]question.Answer{
				{Text: "Answer 1", IsCorrect: true},
				{Text: "Answer 2", IsCorrect: false},
				{Text: "Answer 3", IsCorrect: false},
			},
		),
		*question.New(question.NextQuestionID(),
			"Scopes in closures",
			"What is the expected output?",
			"Some other pseudocode here....",
			[]question.Answer{
				{Text: "aab", IsCorrect: true},
				{Text: "abb", IsCorrect: false},
				{Text: "aba", IsCorrect: false},
			},
		),
	}

	BeforeEach(func() {
		app := webapp.NewWebApp("..", questions)
		ts = httptest.NewServer(app.Mux)
	})

	AfterEach(func() {
		ts.Close()
	})

	Context("/questions", func() {
		It("lists all available questions", func() {
			resp, err := http.Get(ts.URL + "/questions")
			bodyStr := mustReadResponse(resp, err)

			Expect(bodyStr).To(ContainSubstring("Lexical scope"))
			Expect(bodyStr).To(ContainSubstring("Scopes in closures"))
		})
	})

	Context("GET /questions/0", func() {
		It("shows the question", func() {
			resp, err := http.Get(ts.URL + "/questions/0")
			bodyStr := mustReadResponse(resp, err)

			Expect(bodyStr).To(ContainSubstring("Lexical scope"))
			Expect(bodyStr).To(ContainSubstring("Does Powershell do X?"))
			Expect(bodyStr).To(ContainSubstring("Some pseudocode here...."))
			Expect(bodyStr).To(ContainSubstring("Answer 1"))
			Expect(bodyStr).To(ContainSubstring("Answer 2"))
			Expect(bodyStr).To(ContainSubstring("Answer 3"))
		})
	})

	Context("POST /questions/0", func() {
		It("shows congrats if the selected answer is correct", func() {
			formValues := map[string][]string{"answerID": {"0"}}
			resp, err := http.PostForm(ts.URL+"/questions/0", formValues)
			bodyStr := mustReadResponse(resp, err)

			Expect(bodyStr).To(ContainSubstring("Congrats"))
		})

		It("shows condolences if the selected answer is incorrect", func() {
			formValues := url.Values{"answerID": {"1"}}
			resp, err := http.PostForm(ts.URL+"/questions/0", formValues)
			bodyStr := mustReadResponse(resp, err)

			Expect(bodyStr).To(ContainSubstring("Sorry"))
		})

		It("doesn't show the list of possible answers", func() {
			formValues := url.Values{"answerID": {"1"}}
			resp, err := http.PostForm(ts.URL+"/questions/0", formValues)
			bodyStr := mustReadResponse(resp, err)

			Expect(bodyStr).ToNot(ContainSubstring("Answer 1"))
		})
	})
})

func mustReadResponse(resp *http.Response, err error) string {
	Expect(err).ToNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))

	body, err := ioutil.ReadAll(resp.Body)
	Expect(err).ToNot(HaveOccurred())
	return string(body)
}
