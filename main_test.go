package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestInvalidMethodType(t *testing.T) {
	request := httptest.NewRequest("GET", "/add", nil)
	response := httptest.NewRecorder()

	handleRequest(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %v, received %v", http.StatusBadRequest, response.Code)
	}
}

func TestPostValidBody(t *testing.T) {
	reqJSON := "{\"posted_date\":\"1488701255\",\"posted_from\":\"test\",\"content\":\"hello from test\"}"
	request := httptest.NewRequest("POST", "/add", strings.NewReader(reqJSON))
	response := httptest.NewRecorder()

	handleRequest(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("Expected status code %v, received %v", http.StatusCreated, response.Code)
	}

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Error(err)
	}

	respBodyJSON := strings.TrimSpace(string(respBody[:]))

	if respBodyJSON != reqJSON {
		t.Fatalf("Expected response body:\n%v\nreceived:\n%v", reqJSON, respBodyJSON)
	}
}
