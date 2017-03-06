package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type httpErrorMsg struct {
	Message string `json:"message"`
}

func isPostMethod(method string) bool {
	return method == http.MethodPost
}

func replyWithError(w http.ResponseWriter, msg string, code int) {
	errJSON, _ := json.Marshal(httpErrorMsg{msg})
	http.Error(w, string(errJSON[:]), code)
}

// WritePost functions should write a TILPost to a file
type WritePost func(*TILPost, string) (string, error)

func handleRequest(writer WritePost) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isPostMethod(r.Method) {
			replyWithError(w, "invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			replyWithError(w, "invalid json", http.StatusBadRequest)
			return
		}

		var post TILPost
		json.Unmarshal(body, &post)
		defer r.Body.Close()

		_, err = writer(&post, "")

		if err != nil {
			replyWithError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(post)
	}
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v [%v]: %v", time.Now(), r.Method, r.URL.Path)
		f(w, r)
	}
}

func auth(secret string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		potentialSecrets, ok := query["secret"]

		if ok && len(potentialSecrets) == 1 && secret == potentialSecrets[0] {
			f(w, r)
		} else {
			replyWithError(w, "incorrect secret", http.StatusUnauthorized)
		}
	}
}

func startServer(port string, secret string) {
	http.HandleFunc("/add", logging(auth(secret, handleRequest(writePost))))
	http.ListenAndServe(port, nil)
}
