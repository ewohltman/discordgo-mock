// Package mockmember provides functionality for creating a mock
// *discordgo.Member.
package mockmember

import "github.com/bwmarrin/discordgo"

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.Member.
type OptionFunc func(member *discordgo.Member)

// New returns a new *discordgo.Member with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) *discordgo.Member {
	member := &discordgo.Member{}

	for _, optionFunc := range optionFuncs {
		optionFunc(member)
	}

	return member
}

// WithUser sets a *discordgo.Member.User to the given user.
func WithUser(user *discordgo.User) OptionFunc {
	return func(member *discordgo.Member) {
		member.User = user
	}
}

// WithGuildID sets a *discordgo.Member.GuildID to the given guildID.
func WithGuildID(guildID string) OptionFunc {
	return func(member *discordgo.Member) {
		member.GuildID = guildID
	}
}

// WithRoles adds the given roles to a *discordgo.Member.Roles.
func WithRoles(roles ...*discordgo.Role) OptionFunc {
	return func(member *discordgo.Member) {
		for _, role := range roles {
			member.Roles = append(member.Roles, role.ID)
		}
	}
}
