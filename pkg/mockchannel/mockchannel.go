package mockchannel

import "github.com/bwmarrin/discordgo"

type OptionFunc func(channel *discordgo.Channel)

func New(optionFuncs ...OptionFunc) *discordgo.Channel {
	channel := &discordgo.Channel{}

	for _, optionFunc := range optionFuncs {
		optionFunc(channel)
	}

	return channel
}

func WithID(id string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.ID = id
	}
}

func WithGuildID(guildID string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.GuildID = guildID
	}
}

func WithName(name string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.Name = name
	}
}

func WithType(channelType discordgo.ChannelType) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.Type = channelType
	}
}

func WithPermissionOverwrites(permissionOverwrites ...*discordgo.PermissionOverwrite) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.PermissionOverwrites = permissionOverwrites
	}
}
