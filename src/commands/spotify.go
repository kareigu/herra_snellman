package commands

import (
	"github.com/bwmarrin/discordgo"
)

var spotifyInfo = discordgo.ApplicationCommand{
	Name:        "spotify",
	Description: "Spotify soittolista",
}

func spotifyCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: `https://spoti.fi/2BixP98`,
		},
	})
}
