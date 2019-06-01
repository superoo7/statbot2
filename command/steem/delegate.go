package steem

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
)

func DelegateCommand(cID string, c chan<- d.DiscordEmbedMessage, args []string) {
	if len(args) >= 2 {
		withDelegator := len(args) >= 3
		var from, to, sp, link string
		if withDelegator {
			from = args[0]
			to = args[1]
			sp = args[2]
		} else {
			from = ""
			to = args[0]
			sp = args[1]
		}
		title := fmt.Sprintf("Delegate to @%s", to)
		if from == "" {
			link = fmt.Sprintf("https://beta.steemconnect.com/sign/delegateVestingShares?delegatee=%s&vesting_shares=%s%%20SP", to, sp)
			from = "me"
		} else {
			link = fmt.Sprintf("https://beta.steemconnect.com/sign/delegateVestingShares?delegator=%s+&delegatee=%s&vesting_shares=%s%%20SP", from, to, sp)
		}
		em := d.GenMultipleEmbed(d.Green, title, []*discord.MessageEmbedField{
			&discord.MessageEmbedField{
				Name:   "From",
				Value:  from,
				Inline: true,
			},
			&discord.MessageEmbedField{
				Name:   "To",
				Value:  to,
				Inline: true,
			},
			&discord.MessageEmbedField{
				Name:   "SP",
				Value:  sp,
				Inline: true,
			},
			&discord.MessageEmbedField{
				Name:   "SteemConnect link to delegate SP",
				Value:  link,
				Inline: false,
			},
		})
		c <- d.DiscordEmbedMessage{CID: cID, Message: em}
	}
}
