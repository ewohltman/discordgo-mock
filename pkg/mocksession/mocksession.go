package mocksession

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type OptionFunc func(session *discordgo.Session)

// New provides a *discordgo.Session instance to be used in unit testing with
// pre-populated initial state.
func New(optionFuncs ...OptionFunc) (*discordgo.Session, error) {
	session := &discordgo.Session{
		Ratelimiter: discordgo.NewRatelimiter(),
	}

	for _, optionFunc := range optionFuncs {
		optionFunc(session)
	}

	return session, nil
}

func WithState(state *discordgo.State) OptionFunc {
	return func(session *discordgo.Session) {
		session.State = state
		session.StateEnabled = true
	}
}

func WithRESTClient(client *http.Client) OptionFunc {
	return func(session *discordgo.Session) {
		session.Client = client
	}
}
