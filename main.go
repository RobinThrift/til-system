package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
	"time"
)

type unixTime struct {
	Time time.Time
}

type TILPost struct {
	PostedDate unixTime `json:"posted_date"`
	PostedFrom string   `json:"posted_from"`
	Content    string   `json:"content"`
}

func (t *unixTime) UnmarshalJSON(buf []byte) error {
	secondsString := strings.Trim(string(buf[:]), `"`)
	seconds, err := strconv.Atoi(secondsString)
	if err != nil {
		return err
	}

	tt := time.Unix(int64(seconds), 0)
	t.Time = tt
	return nil
}

func (t unixTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(t.Time.Unix(), 10) + `"`), nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "{}", http.StatusBadRequest)
	}

	body, _ := ioutil.ReadAll(r.Body)

	var post TILPost
	json.Unmarshal(body, &post)
	defer r.Body.Close()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(post)
}

func main() {
}
