package main

import "github.com/joho/godotenv"

// Set environment
func setEnv() {
	err := godotenv.Load()

	// TODO: Add error treatment
	if err != nil {
		panic(err)
	}
}

func main() {
	setEnv()
}
