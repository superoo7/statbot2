package command

import (
	"fmt"
	"strconv"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/superoo7/statbot2/config"
	d "github.com/superoo7/statbot2/discord"
	"github.com/superoo7/statbot2/http"
)

// %convert 1 steem usd
func ConvertCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage, args []string) {
	if len(args) >= 3 {
		// Load CoinList
		err := LoadCoinList()
		if err != nil {
			em := d.GenSimpleEmbed(d.Red, "CoinGecko API cannot be connected.")
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}

		// Get amount (args[0])
		a, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			em := d.GenErrorMessage(fmt.Sprintf("`%s` is not a valid number", args[0]))
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}
		amount := float32(a)

		// Get coin1 and coin2 (args[1], args[2])
		c1 := strings.ToLower(args[1])
		c2 := strings.ToLower(args[2])

		var coin1, coin2 string

		// Setup Coin1
		exit, cc1 := IsCoinInList(c1)
		if exit {
			coin1 = c1
		} else {
			coin1 = cc1.ID
		}
		// Setup Coin2
		exit, cc2 := IsCoinInList(c2)
		if exit {
			coin2 = c2
		} else {
			coin2 = cc2.ID
		}

		isCoin1Supported := contains(config.SupportedCurrencies, coin1)
		isCoin2Supported := contains(config.SupportedCurrencies, coin2)
		if isCoin1Supported {
			convertedPriceCoin2, err := http.CG.SimpleSinglePrice(coin2, coin1)
			if err != nil {
				em := d.GenErrorMessage(fmt.Sprintf("%s is not supported", c2))
				c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
				return
			}
			coin2p := convertedPriceCoin2.MarketPrice
			em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("%f %s <=> %f %s", amount, c1, coin2p*amount, c2))
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		} else if isCoin2Supported {
			convertedPriceCoin1, err := http.CG.SimpleSinglePrice(coin1, coin2)
			if err != nil {
				em := d.GenErrorMessage(fmt.Sprintf("%s is not supported", c1))
				c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
				return
			}
			coin1p := convertedPriceCoin1.MarketPrice
			em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("%f %s <=> %f %s", amount, c1, amount/coin1p, c2))
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		} else {
			// maybe crypto -> crypto
			convertedPriceCoin1, err := http.CG.SimpleSinglePrice(coin1, "usd")
			if err != nil {
				em := d.GenErrorMessage(fmt.Sprintf("%s is not supported", c1))
				c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
				return
			}
			p1 := convertedPriceCoin1.MarketPrice
			convertedPriceCoin2, err := http.CG.SimpleSinglePrice(coin2, "usd")
			if err != nil {
				em := d.GenErrorMessage(fmt.Sprintf("%s is not supported", c1))
				c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
				return
			}
			p2 := convertedPriceCoin2.MarketPrice
			em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("%f %s <=> %f %s", amount, c1, amount*p1/p2, c2))
			c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			return
		}
	} else {
		em := d.GenErrorMessage("Invalid format, please try `%convert <amount> <crypto/fiat> <crypto/fiat>`")
		c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
		return
	}
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
