package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type execCommand func(string, ...string) error

func osExec(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	_, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	return nil
}

func gitClone(cmd execCommand, repoURL string, targetDir string) error {
	return cmd("git", "clone", repoURL, targetDir)
}

func removeGitRepo(path string) error {
	return os.RemoveAll(path)
}

func gitAdd(cmd execCommand, gitDir string, filepath string) error {
	return cmd("git", "-C", gitDir, "add", filepath)
}

func gitCommit(cmd execCommand, gitDir string, message string) error {
	return cmd("git", "-C", gitDir, "commit", "-m", message)
}

func gitPush(cmd execCommand, gitDir string) error {
	return cmd("git", "-C", gitDir, "push")
}

func addAndCommitPost(cmd execCommand, gitDir string, filepath string) error {
	err := gitAdd(cmd, gitDir, filepath)
	if err != nil {
		return err
	}

	title := strings.Replace(path.Base(filepath), path.Ext(filepath), "", 1)
	err = gitCommit(cmd, gitDir, fmt.Sprintf("post(%v): add TIL post %v", title, title))
	if err != nil {
		return err
	}

	return nil
}
