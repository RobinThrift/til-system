package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func writerNOOP(p *TILPost, root string) (string, error) {
	return "", nil
}

func TestInvalidMethodType(t *testing.T) {
	request := httptest.NewRequest("GET", "/add", nil)
	response := httptest.NewRecorder()

	handleRequest(writerNOOP, "")(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected status code %v, received %v", http.StatusBadRequest, response.Code)
	}

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Error(err)
	}

	expectedJSON := "{\"message\":\"invalid request method\"}"
	respBodyJSON := strings.TrimSpace(string(respBody[:]))

	if respBodyJSON != expectedJSON {
		t.Fatalf("Expected response body:\n%v\nreceived:\n%v", expectedJSON, respBodyJSON)
	}
}

func TestPostValidBody(t *testing.T) {
	reqJSON := "{\"posted_date\":\"1488701255\",\"posted_from\":\"test\",\"content\":\"hello from test\"}"
	request := httptest.NewRequest("POST", "/add", strings.NewReader(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handleRequest(writerNOOP, "")(response, request)

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

func TestAuthFail(t *testing.T) {
	secret := "totally_secret"
	handler := auth(secret, func(w http.ResponseWriter, r *http.Request) {
		w.Write(nil)
	})

	request := httptest.NewRequest("GET", "/add?secret=totally_wrong", nil)
	response := httptest.NewRecorder()

	handler(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status code %v, received %v", http.StatusOK, response.Code)
	}
}

func TestAuthPass(t *testing.T) {
	secret := "totally_secret"
	handler := auth(secret, func(w http.ResponseWriter, r *http.Request) {
		w.Write(nil)
	})

	request := httptest.NewRequest("GET", fmt.Sprintf("/add?secret=%v", secret), nil)
	response := httptest.NewRecorder()

	handler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Expected status code %v, received %v", http.StatusOK, response.Code)
	}
}
