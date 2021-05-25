package mockguild

import "github.com/bwmarrin/discordgo"

type OptionFunc func(guild *discordgo.Guild)

func New(optionFuncs ...OptionFunc) *discordgo.Guild {
	guild := &discordgo.Guild{}

	for _, optionFunc := range optionFuncs {
		optionFunc(guild)
	}

	return guild
}

func WithID(id string) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.ID = id
	}
}

func WithName(name string) OptionFunc {
	return func(guild *discordgo.Guild) {
		guild.Name = name
	}
}

func WithRoles(roles ...*discordgo.Role) OptionFunc {
	return func(guild *discordgo.Guild) {
		for _, role := range roles {
			guild.Roles = append(guild.Roles, role)
		}
	}
}

func WithChannels(channels ...*discordgo.Channel) OptionFunc {
	return func(guild *discordgo.Guild) {
		for _, channel := range channels {
			guild.Channels = append(guild.Channels, channel)
		}
	}
}

func WithMembers(members ...*discordgo.Member) OptionFunc {
	return func(guild *discordgo.Guild) {
		for _, member := range members {
			guild.Members = append(guild.Members, member)
			guild.MemberCount++
		}
	}
}
