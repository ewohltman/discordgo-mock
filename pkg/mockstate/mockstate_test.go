package mockstate_test

import (
	"testing"

	"github.com/ewohltman/discordgo-mock/pkg/mockconstants"
	"github.com/ewohltman/discordgo-mock/pkg/mockguild"
	"github.com/ewohltman/discordgo-mock/pkg/mockstate"
	"github.com/ewohltman/discordgo-mock/pkg/mockuser"
)

func TestNew(t *testing.T) {
	_, err := mockstate.New(
		mockstate.WithUser(mockuser.New(
			mockuser.WithID("Bot"),
			mockuser.WithUsername("Bot"),
			mockuser.IsBot(true),
		)),
		mockstate.WithGuilds(mockguild.New(
			mockguild.WithID(mockconstants.TestGuild),
		)),
	)
	if err != nil {
		t.Fatal(err)
	}
}
