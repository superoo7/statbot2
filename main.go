package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"

	"github.com/superoo7/statbot2/command"
	"github.com/superoo7/statbot2/command/steem"
	"github.com/superoo7/statbot2/config"
	d "github.com/superoo7/statbot2/discord"
)

func main() {
	env := os.Getenv("ENV")
	bot := d.Discord

	// Register the messageCreate func as a callback for MessageCreate events.
	bot.AddHandlerOnce(botReady)
	bot.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
		messageCreate(s, m, d.DiscordEmbedMessageChannel, d.DiscordMessageChannel)
	})
	err := bot.Open()

	go d.ProcessEmbedMessage(d.DiscordEmbedMessageChannel)
	go d.ProcessMessage(d.DiscordMessageChannel)

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
	go d.UpdateSession(s)
	fmt.Println("Bot is running.")
	s.UpdateStatus(0, "Statbot V2 %help to get started")
}

func messageCreate(s *discord.Session, m *discord.MessageCreate, emc chan<- d.DiscordEmbedMessage, mc chan<- d.DiscordMessage) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// filter whitelist channel
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

	// Update session struct
	go d.UpdateSession(s)

	trigger := string(m.Content[0])
	args := strings.Fields(m.Content[1:])

	if trigger == "$" {
		if len(args) < 1 {
			return
		}
		coin := args[0]
		command.PriceCommand(coin, m, emc)
	} else if trigger == "%" {
		if len(args) < 1 {
			return
		}
		switch args[0] {
		case "p", "price":
			if len(args) >= 2 {
				coin := args[1]
				command.PriceCommand(coin, m, emc)
			} else {
				em := d.GenErrorMessage("Invalid command, try `%price <coin>`")
				emc <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: em}
			}
			break
		case "ping":
			command.PingCommand(m, emc)
			break
		case "discord":
			msg := d.GenSimpleEmbed(d.Blue, "Join our discord channel")
			emc <- d.DiscordEmbedMessage{CID: m.ChannelID, Message: msg}
			mc <- d.DiscordMessage{CID: m.ChannelID, Message: "https://discord.gg/J99vTUS"}
			break
		case "h", "help":
			emc <- d.DiscordEmbedMessage{
				CID: m.ChannelID,
				Message: d.GenMultipleEmbed(
					d.Blue,
					fmt.Sprintf("Help Message (%s)", config.Version),
					[]*discord.MessageEmbedField{
						&discord.MessageEmbedField{
							Name:   "`$<coin>` , `%price <coin>`, `%p <coin>`",
							Value:  "for checking cryptocurrency price",
							Inline: false,
						},
						&discord.MessageEmbedField{
							Name:   "`%discord`",
							Value:  "to join our discord!",
							Inline: false,
						},
						&discord.MessageEmbedField{
							Name:   "`$help` , `%h`",
							Value:  "for help",
							Inline: false,
						},
					},
				),
			}
			break
		case "s", "steem":
			steem.SteemCommand(m.ChannelID, args, emc, mc)
			break
		case "convert", "s/sbd", "sbd/s", "delegate", "bugs", "bug", "hunt", "steemhunt":
			emc <- d.DiscordEmbedMessage{
				CID:     m.ChannelID,
				Message: d.GenErrorMessage("Command are still Work In Progress (WIP) in V2, please wait for the update"),
			}
			break
		default:
			emc <- d.DiscordEmbedMessage{
				CID:     m.ChannelID,
				Message: d.GenErrorMessage("Invalid command, Try `%help` to get started"),
			}
			break
		}
	}
}
