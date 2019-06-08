package command

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"

	d "github.com/superoo7/statbot2/discord"
	"github.com/superoo7/statbot2/http"
)

// PriceCommand `%price <coin>` | `$<coin>` to query price of a certain crypto
func PriceCommand(coin string, m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	err := LoadCoinList()

	if err != nil {
		em := d.GenSimpleEmbed(d.Red, "CoinGecko API cannot be connected.")
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}

	coin = strings.ToLower(coin)

	exit, cc := IsCoinInList(coin)
	if exit {
		em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", coin))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}

	price, err := http.CG.SimpleSinglePrice(cc.ID, "usd")
	if err != nil {
		em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", cc.ID))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
	} else {
		// em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("**%s** (%s) is worth %f %s", cc.Name, cc.ID, price.MarketPrice, price.Currency))
		em := d.GenMultipleEmbed(d.Green, fmt.Sprintf("%s -> USD", cc.Name), []*discord.MessageEmbedField{
			&discord.MessageEmbedField{
				Name:   fmt.Sprintf("**%s** (%s) is worth **%f %s**", cc.Name, cc.ID, price.MarketPrice, strings.ToUpper(price.Currency)),
				Value:  fmt.Sprintf("More details at https://www.coingecko.com/coins/%s?ref=superoo7", cc.ID),
				Inline: false,
			},
		})
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
	}
}
