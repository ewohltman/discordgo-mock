package mockmember

import "github.com/bwmarrin/discordgo"

type OptionFunc func(member *discordgo.Member)

func New(optionFuncs ...OptionFunc) *discordgo.Member {
	member := &discordgo.Member{}

	for _, optionFunc := range optionFuncs {
		optionFunc(member)
	}

	return member
}

func WithUser(user *discordgo.User) OptionFunc {
	return func(member *discordgo.Member) {
		member.User = user
	}
}

func WithGuildID(guildID string) OptionFunc {
	return func(member *discordgo.Member) {
		member.GuildID = guildID
	}
}

func WithRoles(roles ...*discordgo.Role) OptionFunc {
	return func(member *discordgo.Member) {
		for _, role := range roles {
			member.Roles = append(member.Roles, role.ID)
		}
	}
}
