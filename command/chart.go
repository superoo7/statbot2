package command

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
	"github.com/superoo7/statbot2/grpc"
)

// ChartCommand `%chart <coin>`
func ChartCommand(coin string, m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	err := LoadCoinList()

	if err != nil {
		em := d.GenSimpleEmbed(d.Red, "CoinGecko API cannot be connected.")
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
	}

	exit, cc := IsCoinInList(coin)
	if exit {
		em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", coin))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}

	r := grpc.GetChart(cc.ID)

	em := d.GenSimpleImageEmbed(d.Blue, fmt.Sprintf("https://s3.ap-southeast-1.amazonaws.com/statbot.superoo7.com/%s", r.Key), fmt.Sprintf("%s Price Chart", cc.Name))
	c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}

}
