package main

import (
	"github.com/joho/godotenv"
	_ "github.com/rafaelbreno/work-at-olist/cmd/logger"
)

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
