package main

import (
	"fmt"
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
	serverGuild := os.Getenv("DISCORD_SERVER_GUILD")

	dSession, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed to create discord session")
		return
	}

	// Open websocket connection to discord and begin listening.
	err = dSession.Open()
	if err != nil {
		log.Fatal("Failed to open discord websocket connection")
		return
	}

	// Register Slash Command with a string argument
	command, err := dSession.ApplicationCommandCreate(dSession.State.User.ID, serverGuild, &discordgo.ApplicationCommand{
		Name:        "fix-social",
		Description: "Attempts to fix a social media link embed",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "url",
				Description: "URL to fix",
				Required:    true,
			},
		},
	})

	if err != nil {
		fmt.Printf("Cannot create slash command: %v\n", err)
		return
	}

	fmt.Printf("Slash Command %s with ID %s created\n", command.Name, command.ID)

	// Interaction handler
	dSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}
		if i.ApplicationCommandData().Name == "fix-social" {
			input := i.ApplicationCommandData().Options[0].StringValue()
			var response string
			if link.IsUrl(input) {
				url, _ := url.Parse(input)

				fixableLink := link.Link{
					URL: url,
				}

				if fixableLink.IsFixableUrl() {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: fixableLink.Fix(),
						},
					})
					response = fixableLink.Fix()
				} else {
					response = fmt.Sprintf("Unable to fix , %s!", input)
				}
			} else {
				response = fmt.Sprintf("%s is not a valid URL", input)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
		}
	})

	dSession.AddHandler(messageCreate)

	defer dSession.Close()

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
