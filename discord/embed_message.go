package discord

import discord "github.com/bwmarrin/discordgo"

// GenMultipleEmbed Generate Embed Message with multiple description
func GenMultipleEmbed(color int, title string, fields []*discord.MessageEmbedField) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author: &discord.MessageEmbedAuthor{},
		Color:  color,
		Fields: fields,
		Title:  title,
	}
}

// GenSimpleEmbed Generate simple embeded message
func GenSimpleEmbed(color int, description string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author:      &discord.MessageEmbedAuthor{},
		Color:       color,
		Description: description,
	}
}

// GenErrorMessage Generate embeded Error Message
func GenErrorMessage(description string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author:      &discord.MessageEmbedAuthor{},
		Color:       Red,
		Description: description,
	}
}
