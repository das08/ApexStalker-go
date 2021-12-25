package models

func GetTrioTierBadge(env *Environments, rank int) string {
	var badge string
	switch {
	case rank < 1200:
		badge = "<:bronze:910106036197797938>"
	case rank < 2800:
		badge = "<:silver:910102394275233832>"
	case rank < 4800:
		badge = "<:gold:910106036051009556>"
	default:
		badge = "<:platinum:910106036055179294>"
	}

	return badge
}

func GetArenaTierBadge(env *Environments, rank int) string {
	var badge string
	switch {
	case rank < 1600:
		badge = "<:bronze:910106036197797938>"
	case rank < 3200:
		badge = "<:silver:910102394275233832>"
	case rank < 4800:
		badge = "<:gold:910106036051009556>"
	default:
		badge = "<:platinum:910106036055179294>"
	}

	return badge
}
