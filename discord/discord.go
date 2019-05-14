package discord

import (
	"fmt"
	"log"
	"sync"

	discordgo "github.com/bwmarrin/discordgo"
	"github.com/superoo7/statbot2/config"
)

// Discord Session
var Discord *discordgo.Session

// Color
const (
	Green int = 0x00ff00
	Red   int = 0xff0000
	Blue  int = 0x0000ff
)

type DiscordEmbedMessage struct {
	CID     string
	Message *discordgo.MessageEmbed
}

type DiscordMessage struct {
	CID     string
	Message string
}

var DiscordEmbedMessageChannel chan DiscordEmbedMessage
var DiscordMessageChannel chan DiscordMessage

var Session *discordgo.Session

var mutex = &sync.Mutex{}

func init() {
	// INIT DISCORD
	fmt.Println("SETTING UP DISCORD...")
	d, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatal("Invalid Discord Token")
	}
	Discord = d

	// CREATE CHANNEL
	DiscordEmbedMessageChannel = make(chan DiscordEmbedMessage)
	DiscordMessageChannel = make(chan DiscordMessage)
}

func UpdateSession(s *discordgo.Session) {
	mutex.Lock()
	Session = s
	mutex.Unlock()
}

func ProcessEmbedMessage(m <-chan DiscordEmbedMessage) {
	for {
		msg := <-m
		Session.ChannelMessageSendEmbed(msg.CID, msg.Message)
	}
}

func ProcessMessage(m <-chan DiscordMessage) {
	for {
		msg := <-m
		Session.ChannelMessageSend(msg.CID, msg.Message)
	}
}
