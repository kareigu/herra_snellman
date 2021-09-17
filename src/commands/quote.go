package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type quoteRes struct {
	Id     string `json:"_id"`
	Text   string
	Number uint64
}

var quoteInfo = discordgo.ApplicationCommand{
	Name:        "quote",
	Description: "Post a quote",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "search",
			Description: "Type in a search term or the quote number to retrieve that quote. $ gets the newest quote",
			Required:    false,
		},
	},
}

func quoteCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := new(quoteRes)
	search_param := ""
	if len(i.Data.Options) > 0 {
		search_param = i.Data.Options[0].StringValue()
	}

	url := fmt.Sprintf("https://misi.mxrr.dev/api/v1/quotes?search=%s", search_param)
	text := ""
	err := getData(url, &quote)
	if err != nil {
		text = "Virhe tapahtui"
		log.Printf("Error getting quote data: %v", err)
	} else {
		text = quote.Text
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: text,
		},
	})
}
