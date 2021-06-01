// Package mockuser provides functionality for creating a mock *discordgo.User.
package mockuser

import "github.com/bwmarrin/discordgo"

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.User.
type OptionFunc func(user *discordgo.User)

// New returns a new *discordgo.User with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) *discordgo.User {
	user := &discordgo.User{}

	for _, optionFunc := range optionFuncs {
		optionFunc(user)
	}

	return user
}

// WithID sets a *discordgo.User.ID to the given id.
func WithID(id string) OptionFunc {
	return func(user *discordgo.User) {
		user.ID = id
	}
}

// WithUsername sets a *discordgo.User.Username to the given username.
func WithUsername(username string) OptionFunc {
	return func(user *discordgo.User) {
		user.Username = username
	}
}

// WithBotFlag sets a *discordgo.User.Bot to the given bot flag.
func WithBotFlag(bot bool) OptionFunc {
	return func(user *discordgo.User) {
		user.Bot = bot
	}
}
