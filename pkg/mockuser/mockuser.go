package mockuser

import "github.com/bwmarrin/discordgo"

type OptionFunc func(user *discordgo.User)

func New(optionFuncs ...OptionFunc) *discordgo.User {
	user := &discordgo.User{}

	for _, optionFunc := range optionFuncs {
		optionFunc(user)
	}

	return user
}

func WithID(id string) OptionFunc {
	return func(user *discordgo.User) {
		user.ID = id
	}
}

func WithUsername(username string) OptionFunc {
	return func(user *discordgo.User) {
		user.Username = username
	}
}

func IsBot(bot bool) OptionFunc {
	return func(user *discordgo.User) {
		user.Bot = bot
	}
}
