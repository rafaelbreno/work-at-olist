package config

import (
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	// TODO: Add error treatment
	if err != nil {
		panic(err)
	}
}
