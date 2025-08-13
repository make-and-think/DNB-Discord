package discordHandlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it'session a good practice.
	if message.Author.ID == session.State.User.ID {
		return
	}
	if len(message.Attachments) == 0 {
		return
	}

	fmt.Println("messageID", message.ID)
	fmt.Println("messageGuildId", message.GuildID)
	fmt.Println(message.Attachments)

	if len(message.Attachments) > 0 {
		for _, attachment := range message.Attachments {
			fmt.Println(attachment.ProxyURL)
			fmt.Println(attachment.URL)
		}
	}

}
