// Package mockrole provides functionality for creating a mock *discordgo.Role.
package mockrole

import "github.com/bwmarrin/discordgo"

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.Role.
type OptionFunc func(*discordgo.Role)

// New returns a new *discordgo.Role with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) *discordgo.Role {
	role := &discordgo.Role{}

	for _, optionFunc := range optionFuncs {
		optionFunc(role)
	}

	return role
}

// WithID sets a *discordgo.Role.ID to the given id.
func WithID(id string) OptionFunc {
	return func(role *discordgo.Role) {
		role.ID = id
	}
}

// WithName sets a *discordgo.Role.Name to the given name.
func WithName(name string) OptionFunc {
	return func(role *discordgo.Role) {
		role.Name = name
	}
}

// WithPermissions sets a *discordgo.Role.Permissions to the given permissions.
func WithPermissions(permissions int64) OptionFunc {
	return func(role *discordgo.Role) {
		role.Permissions = permissions
	}
}
