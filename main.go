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
	basePath := getenv("TIL_BASE_PATH", "_tils")

	os.MkdirAll(basePath, os.FileMode(0755))

	startServer(port, secret, basePath)
}
