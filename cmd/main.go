package main

import (
	"DNB-Discord/internal/config"
	"DNB-Discord/internal/discordHandlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_, err := config.LoadConfig("configs/.secret.toml")
	if err != nil {
		return
	}
	_, err = config.LoadConfig("configs/settings.toml")
	if err != nil {
		return
	}

	// Create a new Discord session using the provided bot token.
	discordGoBot, err := discordgo.New("Bot " + config.Values.Discord.Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	discordGoBot.AddHandler(discordHandlers.Ready)
	discordGoBot.AddHandler(discordHandlers.MessageCreate)

	discordGoBot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Open the websocket and begin listening.
	err = discordGoBot.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discordGoBot.Close()
}
