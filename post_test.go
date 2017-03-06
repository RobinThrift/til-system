package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func getTestPost() *TILPost {
	testTime, err := time.Parse(time.RFC822, "05 Mar 17 00:00 GMT")
	if err != nil {
		panic(err)
	}

	return &TILPost{
		Content: "Go Is Awesome\nTIL that Go is awesome!",
		PostedDate: unixTime{testTime},
		PostedFrom: "",
	}
}

func TestLoadTemplate(t *testing.T) {
	tmpl, err := loadTemplate("post-template")

	if err != nil {
		t.Error(err)
	}

	if tmpl == nil {
		t.Fail()
	}
}

func TestRenderPost(t *testing.T) {
	fixture, _ := ioutil.ReadFile("./fixtures/post-template")

	testPost := getTestPost()

	res, err := renderPost(testPost)

	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(res.Bytes(), fixture) != 0 {
		t.Fatalf("Expected: \n%v\ngot:\n%v", string(fixture), res)
	}
}

func TestWritePost(t *testing.T) {
	fixture, _ := ioutil.ReadFile("./fixtures/post-template")

	testPost := getTestPost()

	rootTmpDir, _ := ioutil.TempDir("", "post_test")
	defer os.RemoveAll(rootTmpDir)

	path, err := writePost(testPost, rootTmpDir)
	if err != nil {
		t.Error(err)
	}

	writtenContents, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(writtenContents, fixture) != 0 {
		t.Fatalf("Expected: \n%v\ngot:\n%v", string(fixture), string(writtenContents))
	}
}
