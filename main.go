package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid authentication scheme format %v\n", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()

		switch data.Name {
		case "ping":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{Content: "pøng!"},
			})
		default:
			log.Printf("Received unknown command %s", data.Name)
		}
	})

	if err := s.Open(); err != nil {
		log.Fatalf("Unable to connect to Discord %v\n", err)
	}
	defer s.Close()

	log.Println("Connection opened, listening for commands")

	select {}
}
