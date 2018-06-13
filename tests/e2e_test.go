package e2e_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
	"github.com/m-kostrzewa/powershell-for-programmers/core"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End2end Tests Suite")
}

var _ = Describe("Quiz", func() {
	var ts *httptest.Server

	questions := []core.Question{
		{
			Title: "Lexical scope",
			Text:  "Does Powershell do X?",
			Body:  "Some pseudocode here....",
			Answers: []core.Answer{
				{Text: "Answer 1", IsCorrect: true},
				{Text: "Answer 2", IsCorrect: false},
				{Text: "Answer 3", IsCorrect: false},
			},
		},
		{
			Title: "Scopes in closures",
			Text:  "What is the expected output?",
			Body:  "Some other pseudocode here....",
			Answers: []core.Answer{
				{Text: "aab", IsCorrect: true},
				{Text: "abb", IsCorrect: false},
				{Text: "aba", IsCorrect: false},
			},
		},
	}

	BeforeEach(func() {
		app := webapp.NewWebApp("../templates", questions)
		ts = httptest.NewServer(app.Mux)
	})

	AfterEach(func() {
		ts.Close()
	})

	Context("/questions", func() {
		It("lists all available questions", func() {
			resp, err := http.Get(ts.URL + "/questions")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			bodyStr := string(body)

			Expect(bodyStr).To(ContainSubstring("Lexical scope"))
			Expect(bodyStr).To(ContainSubstring("Scopes in closures"))
		})
	})

	Context("/questions/0", func() {
		It("shows the question", func() {
			resp, err := http.Get(ts.URL + "/questions/0")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			bodyStr := string(body)

			Expect(bodyStr).To(ContainSubstring("Lexical scope"))
			Expect(bodyStr).To(ContainSubstring("Does Powershell do X?"))
			Expect(bodyStr).To(ContainSubstring("Some pseudocode here...."))
			Expect(bodyStr).To(ContainSubstring("Answer 1"))
			Expect(bodyStr).To(ContainSubstring("Answer 2"))
			Expect(bodyStr).To(ContainSubstring("Answer 3"))
		})
	})

	Context("/answer/0", func() {
		It("shows congrats if the selected answer is correct", func() {
			values := map[string]int{"answerID": 0}
			formContent, _ := json.Marshal(values)
			resp, err := http.Post(ts.URL+"/answer/0", "application/json", bytes.NewBuffer(formContent))

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			bodyStr := string(body)

			Expect(bodyStr).To(ContainSubstring("Congrats"))
		})

		It("shows condolences if the selected answer is incorrect", func() {
			values := map[string]int{"answerID": 1}
			formContent, _ := json.Marshal(values)
			resp, err := http.Post(ts.URL+"/answer/0", "application/json", bytes.NewBuffer(formContent))

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			bodyStr := string(body)

			Expect(bodyStr).To(ContainSubstring("Sorry"))
		})
	})
})
