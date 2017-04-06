package main

import (
	"reflect"
	"testing"
)

func TestGitClone(t *testing.T) {
	called := false
	cmd := func(name string, args ...string) error {
		called = true
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		expectedArgs := []string{
			"clone",
			"https://github.com/RobinThrift/RobinThrift.com",
			"./RobinThrift.com",
		}

		if !reflect.DeepEqual(args, expectedArgs) {
			t.Fatalf("argument mismatch %s", args)
		}

		return nil
	}

	gitClone(
		cmd,
		"https://github.com/RobinThrift/RobinThrift.com",
		"./RobinThrift.com",
	)

	if !called {
		t.Errorf("cmd function wasn't called :(")
	}
}

func TestGitAdd(t *testing.T) {
	called := false
	cmd := func(name string, args ...string) error {
		called = true
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		expectedArgs := []string{
			"-C",
			"test",
			"add",
			"test.md",
		}

		if !reflect.DeepEqual(args, expectedArgs) {
			t.Fatalf("argument mismatch %s", args)
		}

		return nil
	}

	gitAdd(cmd, "test", "test.md")

	if !called {
		t.Errorf("cmd function wasn't called :(")
	}
}

func TestGitCommit(t *testing.T) {
	called := false
	cmd := func(name string, args ...string) error {
		called = true
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		expectedArgs := []string{
			"-C",
			"test",
			"commit",
			"-m",
			"post(test): add test.md",
		}

		if !reflect.DeepEqual(args, expectedArgs) {
			t.Fatalf("argument mismatch %s", args)
		}

		return nil
	}

	gitCommit(cmd, "test", "post(test): add test.md")

	if !called {
		t.Errorf("cmd function wasn't called :(")
	}
}

func TestGitPush(t *testing.T) {
	called := false
	cmd := func(name string, args ...string) error {
		called = true
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		expectedArgs := []string{
			"-C",
			"test",
			"push",
		}

		if !reflect.DeepEqual(args, expectedArgs) {
			t.Fatalf("argument mismatch %s", args)
		}

		return nil
	}

	gitPush(cmd, "test")

	if !called {
		t.Errorf("cmd function wasn't called :(")
	}
}

func TestAddAndCommitPost(t *testing.T) {
	calledTimes := 0
	cmd := func(name string, args ...string) error {
		calledTimes++
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		if calledTimes == 2 && args[4] != "post(testing): add TIL post testing" {
			t.Fatalf("Wrong commit message: '%v'", args[2])
		}

		return nil
	}

	addAndCommitPost(cmd, "test", "here/testing.md")

	if calledTimes != 2 {
		t.Errorf("cmd function not called twice")
	}
}
