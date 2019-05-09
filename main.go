package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"

	"github.com/superoo7/statbot2/coingecko"
	"github.com/superoo7/statbot2/config"
	d "github.com/superoo7/statbot2/discord"
)

func main() {
	env := os.Getenv("ENV")
	bot := d.Discord

	// Register the messageCreate func as a callback for MessageCreate events.
	bot.AddHandlerOnce(botReady)
	bot.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
		messageCreate(s, m, d.DiscordEmbedMessageChannel)
	})
	err := bot.Open()

	go d.ProcessEmbedMessage(d.DiscordEmbedMessageChannel)

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	if env == "development" {
		fmt.Println("Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		// Cleanly close down the Discord session.
		bot.Close()
	} else {
		defer bot.Close()
		<-make(chan struct{})
	}
}

func botReady(s *discord.Session, r *discord.Ready) {
	d.UpdateSession(s)
	fmt.Println("Bot is running.")
	s.UpdateStatus(0, "Statbot V2 %help to get started")
}

func messageCreate(s *discord.Session, m *discord.MessageCreate, c chan<- d.DiscordEmbedMessage) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check whitelist
	exit := true
	for _, cid := range config.Whitelist {
		if m.ChannelID == cid {
			exit = false
			break
		}
	}
	if exit {
		return
	}

	d.UpdateSession(s)

	trigger := string(m.Content[0])
	args := strings.Fields(m.Content[1:])

	if trigger == "$" {
		if len(args) < 1 {
			return
		}
		coin := args[0]
		price, err := coingecko.CG.SimpleSinglePrice(coin, "usd")
		if err != nil {
			em := d.GenSimpleEmbed(d.Red, fmt.Sprintf("%s not found", coin))
			msg := d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			c <- msg
		} else {
			em := d.GenSimpleEmbed(d.Green, fmt.Sprintf("%s is worth %f %s", price.ID, price.MarketPrice, price.Currency))
			msg := d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			c <- msg
		}
	} else if trigger == "%" {
		if len(args) < 1 {
			return
		}
		switch args[0] {
		case "help":
			msg := "`$coin` - for checking cryptocurrency price"
			s.ChannelMessageSendEmbed(m.ChannelID, d.GenSimpleEmbed(d.Blue, msg))
			break
		default:
			break
		}
	}
}
