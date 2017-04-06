package main

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

type execCommand func(string, ...string) error

func osExec(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func gitClone(cmd execCommand, repoURL string, targetDir string) error {
	return cmd("git", "clone", repoURL, targetDir)
}

func removeGitRepo(cmd execCommand, targetDir string) error {
	return cmd("rm", "-rf", targetDir)
}

func gitAdd(cmd execCommand, filepath string) error {
	return cmd("git", "add", filepath)
}

func gitCommit(cmd execCommand, message string) error {
	return cmd("git", "commit", "-m", message)
}

func gitPush(cmd execCommand) error {
	return cmd("git", "push")
}

func addAndCommitPost(cmd execCommand, filepath string) error {
	err := gitAdd(cmd, filepath)
	if err != nil {
		return err
	}

	title := strings.Replace(path.Base(filepath), path.Ext(filepath), "", 1)
	err = gitCommit(cmd, fmt.Sprintf("post(%v): add TIL post %v", title, title))
	if err != nil {
		return err
	}

	return nil
}
