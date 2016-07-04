package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHandler struct{}

func (handler *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(getAnswer))
	} else if r.Method == "POST" {
		first := r.FormValue("first")
		second := r.FormValue("second")
		answer := first + "+" + second
		w.Write([]byte(answer))
	}
}

const getAnswer = "getAnswer"
const postAnswer = "1+2"

func TestGetRequest(t *testing.T) {
	handler := TestHandler{}
	server := httptest.NewServer(&handler)
	defer server.Close()

	data := GetRequest(server.URL, nil)
	assert.Equal(t, getAnswer, string(data))
}

func TestPostRequest(t *testing.T) {
	handler := TestHandler{}
	server := httptest.NewServer(&handler)
	defer server.Close()

	data := make(map[string]string)
	data["first"] = "1"
	data["second"] = "2"
	answer := PostRequest(server.URL, data)
	assert.Equal(t, postAnswer, string(answer))
}
