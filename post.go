package main

import (
	"strconv"
	"strings"
	"time"
)

type WritePost func(*TILPost, string) error

type unixTime struct {
	Time time.Time
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

type TILPost struct {
	PostedDate unixTime `json:"posted_date"`
	PostedFrom string   `json:"posted_from"`
	Content    string   `json:"content"`
}

func writePost(post *TILPost, root string) error {
	return nil
}
