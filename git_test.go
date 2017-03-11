package main

import "testing"

func TestGitAdd(t *testing.T) {
	called := false
	cmd := func(name string, args ...string) error {
		called = true
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		if args[0] != "add" {
			t.Fatalf("first argument wasn't 'git' :O")
		}

		if args[1] != "test.md" {
			t.Fatalf("second argument wasn't 'test.md'")
		}

		return nil
	}

	gitAdd(cmd, "test.md")

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

		if args[0] != "commit" {
			t.Fatalf("first argument wasn't 'git' :O")
		}

		if args[1] != "-m" {
			t.Fatalf("second argument wasn't '-m'")
		}

		if args[2] != "post(test): add test.md" {
			t.Fatalf("third argument wasn't 'post(test): add test.md'")
		}

		return nil
	}

	gitCommit(cmd, "post(test): add test.md")

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

		if args[0] != "push" {
			t.Fatalf("first argument wasn't 'git' :O")
		}

		return nil
	}

	gitPush(cmd)

	if !called {
		t.Errorf("cmd function wasn't called :(")
	}
}

func TestAddAndCommitPost(t *testing.T) {
	calledTimes := 0
	cmd := func(name string, args ...string) error {
		calledTimes += 1
		if name != "git" {
			t.Fatalf("command name wasn't 'git' :O")
		}

		if calledTimes == 2 && args[2] != "post(testing): add TIL post testing" {
			t.Fatalf("Wrong commit message: '%v'", args[2])
		}

		return nil
	}

	addAndCommitPost(cmd, "here/testing.md")

	if calledTimes != 2 {
		t.Errorf("cmd function not called twice")
	}
}
