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
	v, _ := time.Parse(time.RFC3339, "2019-11-06T00:00:00+08:00")
	t := getTimeRemaining(v)
	var msg string
	if t.d < -5 {
		msg = "Steemfest is over, please wait for next year :)\nCheckout https://steemfest.com/ "
	} else if t.t < 0 {
		msg = fmt.Sprintf("Steemfest 2019 is live (day %d)! \nCheckout https://steemfest.com/", (-t.d)+1)
	} else {
		msg = fmt.Sprintf("Steemfest is coming in %d days %d hours %d minutes and %d seconds! \nCheckout https://steemfest.com/", t.d, t.h, t.m, t.s)
	}
	c <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: d.GenSimpleEmbed(d.Blue, msg)}
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
