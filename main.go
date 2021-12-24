package main

import (
	"apexstalker-go/models"
	"fmt"
	"time"
)

func compare(old models.UserData, new models.Stats) (bool, *[]models.DiscordField, *models.UserDataDetail) {
	timestamp := time.Now().Unix()
	hasUpdate := false
	messageField := []models.DiscordField{}

	messageField = append(messageField, models.DiscordField{Name: "テスト配信", Value: "Hello from Go"})

	level := int(new.Data.Segments[0].Stats.Level.Val)
	trioRank := int(new.Data.Segments[0].Stats.Rank_score.Val)
	arenaRank := int(new.Data.Segments[0].Stats.Arena_Score.Val)
	if timestamp > int64(old.Last_update) && int(new.Data.Segments[0].Stats.Level.Val) > old.Stats.Level {
		hasUpdate = true
		messageField = append(messageField, models.DiscordField{Name: "レベル", Value: fmt.Sprint(old.Stats.Level) + "→" + fmt.Sprint(level) + ":laughing:", Inline: false})
	}
	if timestamp > int64(old.Last_update) && int(new.Data.Segments[0].Stats.Rank_score.Val) != old.Stats.Trio_rank {
		hasUpdate = true
		messageField = append(messageField, models.DiscordField{Name: "トリオRank", Value: models.GetTrioTierBadge(old.Stats.Trio_rank) + fmt.Sprint(old.Stats.Trio_rank) + "→" + models.GetTrioTierBadge(trioRank) + fmt.Sprint(trioRank), Inline: false})
	}
	if timestamp > int64(old.Last_update) && int(new.Data.Segments[0].Stats.Arena_Score.Val) != old.Stats.Arena_rank {
		hasUpdate = true
		messageField = append(messageField, models.DiscordField{Name: "アリーナRank", Value: models.GetArenaTierBadge(old.Stats.Arena_rank) + fmt.Sprint(old.Stats.Arena_rank) + "→" + models.GetArenaTierBadge(arenaRank) + fmt.Sprint(arenaRank), Inline: false})
	}

	userDataDetail := models.UserDataDetail{Level: level, Trio_rank: trioRank, Arena_rank: arenaRank}

	return hasUpdate, &messageField, &userDataDetail
}

func main() {
	// Load environment values
	envs := models.LoadEnv()

	db := models.Connect()
	defer db.Close()
	userList := models.GetPlayerData(db)

	for _, v := range userList {
		fmt.Printf("Data: %+v\n", v)
		userStats, err := models.GetApexStats(envs.APEX_API_ENDPOINT, envs.APEX_API_KEY, v.Platform, v.Uid)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", userStats.Data)
		// models.UpsertPlayerData(db, )
		hasUpdate, messageField, userDataDetail := compare(v, *userStats)
		models.UpdatePlayerData(db, v.Uid, *userDataDetail)

		if hasUpdate {
			msgObj := new(models.DiscordWebhook)
			msgObj.UserName = "Go BOT"
			msgObj.AvatarURL = "https://pbs.twimg.com/profile_images/1108370004590772224/hEX1gucp_400x400.jpg"
			msgObj.Embeds = []models.DiscordEmbed{
				{
					Title:  "\U0001F38A" + v.Uid + "の戦績変化\U0001F389",
					Color:  0x550000,
					Fields: *messageField,
				},
			}

			// models.SendMessage(envs.DISCORD_ENDPOINT, msgObj)
		}
	}

}
