package main

import (
	"encoding/json"
	"errors"
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

func isJSONRequest(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func addJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func replyWithError(w http.ResponseWriter, msg string, code int) {
	addJSONContentType(w)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(httpErrorMsg{msg})
	if err != nil {
		panic(err)
	}
}

func validatePost(post *TILPost) error {
	if len(post.Content) == 0 {
		return errors.New("Please provide content")
	}

	return nil
}

// WritePost functions should write a TILPost to a file
type WritePost func(*TILPost, string) (string, error)

func handleRequest(writer WritePost, basePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isPostMethod(r.Method) {
			replyWithError(w, "invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if !isJSONRequest(r) {
			replyWithError(w, "invalid content type", http.StatusBadRequest)
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

		err = validatePost(&post)
		if err != nil {
			replyWithError(w, "invalid json", http.StatusBadRequest)
			return
		}

		_, err = writer(&post, basePath)

		if err != nil {
			replyWithError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		addJSONContentType(w)
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

func startServer(port string, secret string, basePath string) {
	http.HandleFunc("/add", logging(auth(secret, handleRequest(writePost, basePath))))
	http.ListenAndServe(":" + port, nil)
}
