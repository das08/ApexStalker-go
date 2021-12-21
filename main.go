package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Environments struct {
	APEX_API_ENDPOINT  string
	APEX_API_KEY       string
	DISCORD_ENDPOINT   string
	TINAX_API_ENDPOINT string
}

type User struct {
	uid        string
	platform   string
	level      int
	trio_rank  int
	arena_rank int
}

// type DiscordMessageObj struct {
// 	Title  string         `json:"title"`
// 	Desc   string         `json:"description"`
// 	URL    string         `json:"url"`
// 	Color  int            `json:"color"`
// 	Image  discordImg     `json:"image"`
// 	Thum   discordImg     `json:"thumbnail"`
// 	Author discordAuthor  `json:"author"`
// 	Fields []discordField `json:"fields"`
// }

// type DiscordMessage struct {
// 	user_name  string `json:"username"`
// 	avatar_url string `json:"avatar_url"`
// 	content    string `json:"content"`
// }

type discordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type discordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type discordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  discordImg     `json:"image"`
	Thum   discordImg     `json:"thumbnail"`
	Author discordAuthor  `json:"author"`
	Fields []discordField `json:"fields"`
}

type discordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []discordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
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

func sendMessage(discord_endpoint string, msgObj *discordWebhook) {
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

	// msgObj := new(DiscordMessage)
	// msgObj.user_name = "Go BOT"
	// msgObj.avatar_url = "https://pbs.twimg.com/profile_images/1108370004590772224/hEX1gucp_400x400.jpg"
	// msgObj.content = "Hello from Go."
	msgObj := &discordWebhook{UserName: "Narumium", Content: "Webhook Test"}
	sendMessage(envs.DISCORD_ENDPOINT, msgObj)
}
