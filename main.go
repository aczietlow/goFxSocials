package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"goFxSocials/link"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")

	dSession, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed to create discord session")
		return
	}

	defer dSession.Close()

	dSession.AddHandler(messageCreate)

	// Open websocket connection to discord and begin listening.
	err = dSession.Open()
	if err != nil {
		log.Fatal("Failed to open discord websocket connection")
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("goFxSocials is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "poke" {
		s.ChannelMessageSend(m.ChannelID, "oooh myyy")
	}

	if link.IsUrl(m.Content) {
		url, _ := url.Parse(m.Content)
		fixableLink := link.Link{
			URL: url,
		}

		if fixableLink.IsFixableUrl() {
			s.ChannelMessageSend(m.ChannelID, "Let me try to fix that for you")
			s.ChannelMessageSend(m.ChannelID, fixableLink.Fix())
		}
	}
}
