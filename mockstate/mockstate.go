// Package mockstate provides functionality for creating a mock
// *discordgo.State.
package mockstate

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/ewohltman/discordgo-mock/mockconstants"
)

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.State.
type OptionFunc func(state *discordgo.State) error

// New returns a new *discordgo.State with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) (*discordgo.State, error) {
	state := discordgo.NewState()

	state.User = &discordgo.User{
		ID:       mockconstants.TestSession,
		Username: mockconstants.TestSession,
		Bot:      true,
	}

	for _, optionFunc := range optionFuncs {
		err := optionFunc(state)
		if err != nil {
			return nil, fmt.Errorf("error applying option func: %w", err)
		}
	}

	return state, nil
}

// WithUser sets a *discordgo.State.User to the given user.
func WithUser(user *discordgo.User) OptionFunc {
	return func(state *discordgo.State) error {
		state.User = user

		return nil
	}
}

// WithGuilds adds the given guilds to a *discordgo.State.
func WithGuilds(guilds ...*discordgo.Guild) OptionFunc {
	return func(state *discordgo.State) error {
		for _, guild := range guilds {
			err := state.GuildAdd(guild)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
