package discord

import discord "github.com/bwmarrin/discordgo"

// GenSimpleEmbed Generate simple embeded message
func GenSimpleEmbed(color int, description string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Author:      &discord.MessageEmbedAuthor{},
		Color:       color,
		Description: description,
	}
}
