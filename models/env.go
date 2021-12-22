package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environments struct {
	APEX_API_ENDPOINT  string
	APEX_API_KEY       string
	DISCORD_ENDPOINT   string
	TINAX_API_ENDPOINT string
}

func LoadEnv() Environments {
	// LOADS .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Err: Loading .env failed.")
	}

	env := new(Environments)

	// Load environment values
	env.APEX_API_ENDPOINT = os.Getenv("API_ENDPOINT")
	env.APEX_API_KEY = os.Getenv("API_KEY")
	env.DISCORD_ENDPOINT = os.Getenv("DISCORD_ENDPOINT")
	env.TINAX_API_ENDPOINT = os.Getenv("TINAX_API")

	return *env
}
