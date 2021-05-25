package mockstate_test

import (
	"testing"

	"github.com/ewohltman/discordgo-mock/mockconstants"
	"github.com/ewohltman/discordgo-mock/mockguild"
	"github.com/ewohltman/discordgo-mock/mockstate"
	"github.com/ewohltman/discordgo-mock/mockuser"
)

func TestNew(t *testing.T) {
	_, err := mockstate.New(
		mockstate.WithUser(mockuser.New(
			mockuser.WithID("Bot"),
			mockuser.WithUsername("Bot"),
			mockuser.WithBotFlag(true),
		)),
		mockstate.WithGuilds(mockguild.New(
			mockguild.WithID(mockconstants.TestGuild),
		)),
	)
	if err != nil {
		t.Fatal(err)
	}
}
