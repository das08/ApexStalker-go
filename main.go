package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"apexstalker-go/models"

	"github.com/joho/godotenv"
)

type Environments struct {
	APEX_API_ENDPOINT  string
	APEX_API_KEY       string
	DISCORD_ENDPOINT   string
	TINAX_API_ENDPOINT string
}

func loadEnv() Environments {
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

func sendMessage(discord_endpoint string, msgObj *models.DiscordWebhook) {
	msgJson, err := json.Marshal(msgObj)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", discord_endpoint, bytes.NewBuffer(msgJson))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	if resp.StatusCode == 204 {
		fmt.Println("Success") //成功
	} else {
		fmt.Printf("%#v\n", resp) //失敗
	}
}

func main() {
	envs := loadEnv()
	fmt.Printf("a %v", envs.APEX_API_ENDPOINT)

	msgObj := new(models.DiscordWebhook)
	msgObj.UserName = "Go BOT"
	msgObj.AvatarURL = "https://pbs.twimg.com/profile_images/1108370004590772224/hEX1gucp_400x400.jpg"
	msgObj.Content = "Hello from Go."

	sendMessage(envs.DISCORD_ENDPOINT, msgObj)
}
