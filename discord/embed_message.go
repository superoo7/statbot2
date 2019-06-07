package discord

import discord "github.com/bwmarrin/discordgo"

var footerMsg = discord.MessageEmbedFooter{
	Text:         "Consider donate to us with \"%donate\" command.",
	IconURL:      "https://wenghan.me/statics/skills/statbot.png",
	ProxyIconURL: "https://wenghan.me/statics/skills/statbot.png",
}

// GenMultipleEmbed Generate Embed Message with multiple description
func GenMultipleEmbed(color int, title string, fields []*discord.MessageEmbedField) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author: &discord.MessageEmbedAuthor{},
		Color:  color,
		Fields: fields,
		Title:  title,
		Footer: &footerMsg,
	}
}

// GenSimpleEmbed Generate simple embeded message
func GenSimpleEmbed(color int, description string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author:      &discord.MessageEmbedAuthor{},
		Color:       color,
		Description: description,
		Footer:      &footerMsg,
	}
}

func GenSimpleImageEmbed(color int, url string, title string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author: &discord.MessageEmbedAuthor{},
		Color:  color,
		Title:  title,
		Image: &discord.MessageEmbedImage{
			URL: url,
		},
		Footer: &footerMsg,
	}
}

// GenErrorMessage Generate embeded Error Message
func GenErrorMessage(description string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author:      &discord.MessageEmbedAuthor{},
		Color:       Red,
		Description: description,
		Footer:      &footerMsg,
	}
}
