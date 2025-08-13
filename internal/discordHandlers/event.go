package discordHandlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func Ready(session *discordgo.Session, event *discordgo.Ready) {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	fmt.Println(session.State.User.ID)

	err := session.UpdateGameStatus(0, "HI THERE")
	if err != nil {
		return
	}

}
