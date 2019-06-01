package command

// import (
// 	"strconv"

// 	discord "github.com/bwmarrin/discordgo"
// 	"github.com/superoo7/statbot2/config"
// 	d "github.com/superoo7/statbot2/discord"
// )

// // %convert 1 steem usd
// func ConverCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage, args []string) {
// 	if len(args) >= 3 {
// 		amount, err := strconv.ParseFloat(args[0], 64)
// 		if err != nil {
// 			// handle error
// 		}
// 		coin1 := args[1]
// 		coin2 := args[2]
// 		isCoin1Supported := contains(config.SupportedCurrencies, coin1)
// 		isCoin2Supported := contains(config.SupportedCurrencies, coin2)
// 		if isCoin1Supported && isCoin2Supported {
// 		} else if isCoin1Supported {
// 		} else if isCoin2Supported {
// 		} else {
// 		}

// 	}
// }

// // Find returns the smallest index i at which x == a[i],
// // or len(a) if there is no such index.
// func find(a []string, x string) int {
// 	for i, n := range a {
// 		if x == n {
// 			return i
// 		}
// 	}
// 	return len(a)
// }

// // Contains tells whether a contains x.
// func contains(a []string, x string) bool {
// 	for _, n := range a {
// 		if x == n {
// 			return true
// 		}
// 	}
// 	return false
// }
