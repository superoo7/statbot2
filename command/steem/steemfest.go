package steem

import (
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
	d "github.com/superoo7/statbot2/discord"
)

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func SteemFestCommand(m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {
	v, _ := time.Parse(time.RFC3339, "2019-11-06T14:00:00+08:00")
	t := getTimeRemaining(v)
	var em *discord.MessageEmbed
	if t.t < 0 {
		em = d.GenSimpleEmbed(d.Blue, "Steemfest is live! \nCheckout https://steemfest.com/")
	} else {
		em = d.GenSimpleEmbed(d.Blue, fmt.Sprintf("Steemfest is coming in %d days %d hours %d minutes and %d seconds! \nCheckout https://steemfest.com/", t.d, t.h, t.m, t.s))
	}
	c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
}

func getTimeRemaining(t time.Time) countdown {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	currentTime := time.Now().In(loc)
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
