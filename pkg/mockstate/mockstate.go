package mockstate

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/discordgo-mock/pkg/mockconstants"
)

type OptionFunc func(state *discordgo.State) error

// New provides a *discordgo.State instance to be used in unit testing.
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

func WithUser(user *discordgo.User) OptionFunc {
	return func(state *discordgo.State) error {
		state.User = user
		return nil
	}
}

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
