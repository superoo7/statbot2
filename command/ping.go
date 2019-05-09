package command

import (
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
)

// PingCommand `%ping`
func PingCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	t, err := time.Parse(time.RFC3339, string(m.Timestamp))
	if err != nil {
		em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("Pong!"))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}
	diff := time.Now().Sub(t)
	em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("Pong! %d ms", int(diff/time.Millisecond)))
	c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
}
