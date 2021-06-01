// Package mockchannel provides functionality for creating a mock
// *discordgo.Channel.
package mockchannel

import "github.com/bwmarrin/discordgo"

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.Channel.
type OptionFunc func(channel *discordgo.Channel)

// New returns a new *discordgo.Channel with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) *discordgo.Channel {
	channel := &discordgo.Channel{}

	for _, optionFunc := range optionFuncs {
		optionFunc(channel)
	}

	return channel
}

// WithID sets a *discordgo.Channel.ID to the given id.
func WithID(id string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.ID = id
	}
}

// WithGuildID sets a *discordgo.Channel.GuildID to the given guildID.
func WithGuildID(guildID string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.GuildID = guildID
	}
}

// WithName sets a *discordgo.Channel.Name to the given name.
func WithName(name string) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.Name = name
	}
}

// WithType sets a *discordgo.Channel.Type to the given channelType.
func WithType(channelType discordgo.ChannelType) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.Type = channelType
	}
}

// WithPermissionOverwrites sets a *discordgo.Channel.PermissionOverwrites to
// the given permissionOverwrites.
func WithPermissionOverwrites(permissionOverwrites ...*discordgo.PermissionOverwrite) OptionFunc {
	return func(channel *discordgo.Channel) {
		channel.PermissionOverwrites = permissionOverwrites
	}
}
