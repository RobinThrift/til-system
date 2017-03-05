package main

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"
)


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
	fixture, err := ioutil.ReadFile("./fixtures/post-template")
	if err != nil {
		t.Fatal(err)
	}

	testTime, err := time.Parse(time.RFC822, "05 Mar 17 00:00 GMT")
	if err != nil {
		t.Fatal(err)
	}

	res, err := renderPost(&TILPost{
		Content: "Go Is Awesome\nTIL that Go is awesome!",
		PostedDate: unixTime{testTime},
		PostedFrom: "",
	})

	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(res.Bytes(), fixture) != 0 {
		t.Fatalf("Expected: \n%v\ngot:\n%v", string(fixture), res)
	}
}
