package main

import (
	"fmt"

	"apexstalker-go/models"
)

func main() {
	// Load environment values
	envs := models.LoadEnv()

	msgObj := new(models.DiscordWebhook)
	msgObj.UserName = "Go BOT"
	msgObj.AvatarURL = "https://pbs.twimg.com/profile_images/1108370004590772224/hEX1gucp_400x400.jpg"
	msgObj.Content = "Hello from Go."

	// models.SendMessage(envs.DISCORD_ENDPOINT, msgObj)
	db := models.Connect()
	defer db.Close()
	userList := models.GetPlayers(db)

	for _, v := range userList {
		fmt.Printf("Data: %+v\n", v)
		userStats, err := models.GetApexStats(envs.APEX_API_ENDPOINT, envs.APEX_API_KEY, v.Platform, v.Uid)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", userStats.Data)
	}
}
