package command

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
	"github.com/superoo7/statbot2/grpc"
)

func DailyCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	r := grpc.GetDailyChart()

	em := d.GenSimpleImageEmbed(d.Blue, fmt.Sprintf("https://s3.ap-southeast-1.amazonaws.com/statbot.superoo7.com/%s", r.Key), "Daily CryptoCurrency Chart")
	c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
}
