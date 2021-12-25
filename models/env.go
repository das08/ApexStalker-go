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
	BRONZE_BADGE       string
	SILVER_BADGE       string
	GOLD_BADGE         string
	PLATINUM_BADGE     string
}

func LoadEnv(debug bool) Environments {
	var err error
	// LOADS .env file
	if debug {
		err = godotenv.Load(".env_local")
	} else {
		err = godotenv.Load(".env_prod")
	}

	if err != nil {
		fmt.Println("Err: Loading .env failed.")
	}

	env := new(Environments)

	// Load environment values
	env.APEX_API_ENDPOINT = os.Getenv("API_ENDPOINT")
	env.APEX_API_KEY = os.Getenv("API_KEY")
	env.DISCORD_ENDPOINT = os.Getenv("DISCORD_ENDPOINT")
	env.TINAX_API_ENDPOINT = os.Getenv("TINAX_API")
	env.BRONZE_BADGE = os.Getenv("BRONZE_BADGE")
	env.SILVER_BADGE = os.Getenv("SILVER_BADGE")
	env.GOLD_BADGE = os.Getenv("GOLD_BADGE")
	env.PLATINUM_BADGE = os.Getenv("PLATINUM_BADGE")

	return *env
}
