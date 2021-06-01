// Package mockguild provides functionality for creating a mock
// *discordgo.Guild.
package mockguild

import "github.com/bwmarrin/discordgo"

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.Guild.
type OptionFunc func(guild *discordgo.Guild)

// New returns a new *discordgo.Guild with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) *discordgo.Guild {
	guild := &discordgo.Guild{}

	for _, optionFunc := range optionFuncs {
		optionFunc(guild)
	}

	return guild
}

// WithID sets a *discordgo.Guild.ID to the given id.
func WithID(id string) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.ID = id
	}
}

// WithName sets a *discordgo.Guild.Name to the given name.
func WithName(name string) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.Name = name
	}
}

// WithRoles adds the given roles to a *discordgo.Guild.
func WithRoles(roles ...*discordgo.Role) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.Roles = append(guild.Roles, roles...)
	}
}

// WithChannels adds the given channels to a *discordgo.Guild.
func WithChannels(channels ...*discordgo.Channel) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.Channels = append(guild.Channels, channels...)
	}
}

// WithMembers adds the given members to a *discordgo.Guild.
func WithMembers(members ...*discordgo.Member) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.Members = append(guild.Members, members...)
		guild.MemberCount = len(guild.Members)
	}
}
