package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Link struct {
	URL      *url.URL
	Hostname string
}

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

	//https://www.instagram.com/reel/C1pPFU9rLym/

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "poke" {
		s.ChannelMessageSend(m.ChannelID, "oooh myyy")
	}

	if isUrl(m.Content) {
		url, _ := url.Parse(m.Content)
		hostname := getHostname(url.Hostname())
		fmt.Println(hostname)
		if isFixableUrl(hostname) {
			s.ChannelMessageSend(m.ChannelID, "Let me try to fix that for you")
			fixableLink := Link{
				URL:      url,
				Hostname: hostname,
			}
			s.ChannelMessageSend(m.ChannelID, fixURL(fixableLink))

		}
	}
}

func isUrl(str string) bool {
	url, err := url.Parse(str)
	if err != nil {
		return false
	}
	if url.Scheme == "" || url.Host == "" {
		return false
	}
	return true
}

func getHostname(hostname string) string {
	// Check if the hostname has www. in it and if so remove it
	substring := strings.Split(hostname, ".")

	if substring[0] == "www" {
		hostname = strings.Join(substring[1:], ".")
	}

	return hostname
}

// Checks if string is a valid url
func isFixableUrl(hostname string) bool {
	switch hostname {
	case "instagram.com", "twitter.com", "x.com", "reddit.com", "tiktok.com":
		return true
	}
	return false
}

func fixURL(link Link) string {
	switch link.Hostname {
	case "instagram.com":
		return fixInstagram(link)
	case "twitter.com":
		return fixTwitter(link)
	case "x.com":
		return fixX(link)
	case "reddit.com":
		return fixReddit(link)
	case "tiktok.com":
		return fixTikTok(link)

	}
	return ""
}

func fixInstagram(link Link) string {
	newURL := link.URL.Scheme + "://dd" + link.Hostname + link.URL.Path
	return newURL
}

func fixTwitter(link Link) string {
	newURL := link.URL.Scheme + "://fx" + link.Hostname + link.URL.Path
	return newURL
}

func fixX(link Link) string {
	newURL := link.URL.Scheme + "://fixv" + link.Hostname + link.URL.Path
	return newURL
}

func fixReddit(link Link) string {
	newURL := link.URL.Scheme + "://vx" + link.Hostname + link.URL.Path
	return newURL
}

func fixTikTok(link Link) string {
	newURL := link.URL.Scheme + "://vx" + link.Hostname + link.URL.Path
	return newURL
}
