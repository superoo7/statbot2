package discord

import (
	"log"
	"os"

	discord "github.com/bwmarrin/discordgo"
)

// Discord Session
var Discord *discord.Session

// Color
const (
	Green int = 0x00ff00
	Red   int = 0xff0000
	Blue  int = 0x0000ff
)

type DiscordEmbedMessage struct {
	CID     string
	Message *discord.MessageEmbed
}

type DiscordMessage struct {
	CID     string
	Message string
}

var DiscordEmbedMessageChannel chan DiscordEmbedMessage
var DiscordMessageChannel chan DiscordMessage

var Session *discord.Session

func init() {
	// INIT DISCORD
	token := os.Getenv("DISCORD_TOKEN")
	d, err := discord.New("Bot " + token)
	if err != nil {
		log.Fatal("Invalid Discord Token")
	}
	Discord = d

	// CREATE CHANNEL
	DiscordEmbedMessageChannel = make(chan DiscordEmbedMessage)
	DiscordMessageChannel = make(chan DiscordMessage)
}

func UpdateSession(s *discord.Session) {
	Session = s
}

func ProcessEmbedMessage(m <-chan DiscordEmbedMessage) {
	for {
		msg := <-m
		Session.ChannelMessageSendEmbed(msg.CID, msg.Message)
	}
}
