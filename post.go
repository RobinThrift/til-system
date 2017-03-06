package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"
)

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

// TILPost is the main struct to describe a TIL item.
// In most cases this will be constructed from the received JSON.
type TILPost struct {
	PostedDate unixTime `json:"posted_date"`
	PostedFrom string   `json:"posted_from"`
	Content    string   `json:"content"`
}

func loadTemplate(templatePath string) (*template.Template, error) {
	tmplString, err := Asset("assets/" + templatePath + ".tmpl")

	if err != nil {
		return nil, err
	}

	tmpl := template.New(templatePath)
	_, err = tmpl.Parse(string(tmplString[:]))

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

type templateData struct {
	Title   string
	Date    string
	Slug    string
	Content string
}

func renderPost(post *TILPost) (*bytes.Buffer, error) {
	tmpl, err := loadTemplate("post-template")

	if err != nil {
		return nil, err
	}

	lines := strings.Split(post.Content, "\n")

	data := templateData{
		Title:   lines[0],
		Content: strings.Join(lines[1:], "\n"),
		Date:    post.PostedDate.Time.Format("2006-01-02"),
		Slug:    strings.Replace(strings.ToLower(lines[0]), " ", "-", -1),
	}

	buff := new(bytes.Buffer)
	err = tmpl.Execute(buff, &data)

	if err != nil {
		return nil, err
	}

	return buff, nil
}

func writePost(post *TILPost, root string) (string, error) {
	fileName := fmt.Sprintf("til-%v.md", post.PostedDate.Time.Format("2006-01-02"))
	filePath := path.Join(root, fileName)

	contents, err := renderPost(post)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filePath, contents.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
