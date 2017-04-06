package main

import "os"

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	port := getenv("TIL_PORT", "3000")
	secret := getenv("TIL_SECRET", "PLEASE_SET_A_SECRET")
	repoURL := getenv("TIL_REPO_URL", "INVALIAD_REPO_URL")
	postDir := getenv("TIL_POST_DIR", "content/til")

	startServer(port, secret, repoURL, postDir)
}
