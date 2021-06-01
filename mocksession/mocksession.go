// Package mocksession provides functionality for creating a mock
// *discordgo.Session.
package mocksession

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// OptionFunc is a function that can be used to apply different options to a
// *discordgo.Session.
type OptionFunc func(session *discordgo.Session)

// New returns a new *discordgo.Session with the given optionFuncs applied.
func New(optionFuncs ...OptionFunc) (*discordgo.Session, error) {
	session := &discordgo.Session{
		Ratelimiter: discordgo.NewRatelimiter(),
	}

	for _, optionFunc := range optionFuncs {
		optionFunc(session)
	}

	return session, nil
}

// WithState sets a *discordgo.Session.State to the given state and sets
// *discordgo.Session.StateEnabled to true.
func WithState(state *discordgo.State) OptionFunc {
	return func(session *discordgo.Session) {
		session.State = state
		session.StateEnabled = true
	}
}

// WithClient sets a *discordgo.Session.Client to the given client.
func WithClient(client *http.Client) OptionFunc {
	return func(session *discordgo.Session) {
		session.Client = client
	}
}
