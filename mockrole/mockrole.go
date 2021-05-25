package mockrole

import "github.com/bwmarrin/discordgo"

type OptionFunc func(*discordgo.Role)

func New(optionFuncs ...OptionFunc) *discordgo.Role {
	role := &discordgo.Role{}

	for _, optionFunc := range optionFuncs {
		optionFunc(role)
	}

	return role
}

func WithID(id string) OptionFunc {
	return func(role *discordgo.Role) {
		role.ID = id
	}
}

func WithName(name string) OptionFunc {
	return func(role *discordgo.Role) {
		role.Name = name
	}
}

func WithPermissions(permissions int64) OptionFunc {
	return func(role *discordgo.Role) {
		role.Permissions = permissions
	}
}
