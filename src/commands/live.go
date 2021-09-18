package commands

import (
	"github.com/bwmarrin/discordgo"
)

var liveInfo = discordgo.ApplicationCommand{
	Name:        "live",
	Description: "Linkki striimiin",
}

func liveCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: `https://www.youtube.com/c/Misikaani/live`,
		},
	})
}
