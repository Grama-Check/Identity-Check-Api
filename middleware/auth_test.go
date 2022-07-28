package middleware

import (
	"net/http"
	"testing"
)

func testAuth(t *testing.T) {

	// creating new request

	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Add("authentication", "bearer ")

}
