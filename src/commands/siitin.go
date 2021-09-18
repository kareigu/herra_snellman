package commands

import (
	"github.com/bwmarrin/discordgo"
)

var siitinInfo = discordgo.ApplicationCommand{
	Name:        "siitin",
	Description: "Sääli on sairautta",
}

func siitinCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: `"Mutta kaikesta huolimatta. Olen kivenkova rasisti, sadisti ja fasisti. Siitä minä nautin. Sääli on sairautta." -Valtakunnanjohtaja Siitoin`,
		},
	})
}
