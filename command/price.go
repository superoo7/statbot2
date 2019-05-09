package command

import (
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
	"github.com/superoo7/go-gecko/v3/types"

	"github.com/superoo7/statbot2/coingecko"
	d "github.com/superoo7/statbot2/discord"
)

var coinlist *types.CoinList
var lastSavedCoinList time.Time

// PriceCommand `%price <coin>` | `$<coin>` to query price of a certain crypto
func PriceCommand(coin string, m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	now := time.Now()
	diff := now.Sub(lastSavedCoinList)
	days := int(diff.Hours() / 24)

	if days > 1 {
		cl, err := coingecko.CG.CoinsList()
		if err != nil {
			em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", coin))
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}
		coinlist = cl
		lastSavedCoinList = now
	}

	var cc types.CoinsListItem

	exit := true
	for _, c := range *coinlist {
		if coin == c.ID || coin == c.Name || coin == c.Symbol {
			exit = false
			cc = c
			break
		}
	}
	if exit {
		em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", coin))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}

	price, err := coingecko.CG.SimpleSinglePrice(cc.ID, "usd")
	if err != nil {
		em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", cc.ID))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
	} else {
		em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("**%s** (%s) is worth %f %s", cc.Name, cc.ID, price.MarketPrice, price.Currency))
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
	}
}
