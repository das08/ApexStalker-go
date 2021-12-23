package main

import (
	"fmt"

	"apexstalker-go/models"
)

func main() {
	envs := models.LoadEnv()
	fmt.Printf("a %v", envs.APEX_API_ENDPOINT)

	msgObj := new(models.DiscordWebhook)
	msgObj.UserName = "Go BOT"
	msgObj.AvatarURL = "https://pbs.twimg.com/profile_images/1108370004590772224/hEX1gucp_400x400.jpg"
	msgObj.Content = "Hello from Go."

	// models.SendMessage(envs.DISCORD_ENDPOINT, msgObj)
	userList := models.GetPlayers()

	for _, v := range userList {
		fmt.Printf("Data: %+v\n", v)
		// models.GetApexStats(envs.APEX_API_ENDPOINT, envs.APEX_API_KEY, v.Platform, v.Uid)
	}

}
