package mocksession_test

import (
	"net/http"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/discordgo-mock/pkg/mockchannel"
	"github.com/ewohltman/discordgo-mock/pkg/mockconstants"
	"github.com/ewohltman/discordgo-mock/pkg/mockguild"
	"github.com/ewohltman/discordgo-mock/pkg/mockmember"
	"github.com/ewohltman/discordgo-mock/pkg/mockrest"
	"github.com/ewohltman/discordgo-mock/pkg/mockrole"
	"github.com/ewohltman/discordgo-mock/pkg/mocksession"
	"github.com/ewohltman/discordgo-mock/pkg/mockstate"
	"github.com/ewohltman/discordgo-mock/pkg/mockuser"
)

func TestNew(t *testing.T) {
	state, err := newState()
	if err != nil {
		t.Fatal(err)
	}

	session, err := mocksession.New(
		mocksession.WithState(state),
		mocksession.WithClient(&http.Client{
			Transport: mockrest.NewTransport(state),
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	guildBefore, err := session.Guild(mockconstants.TestGuild)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Name: %s", guildBefore.Name)
	t.Logf("Channels: %d", len(guildBefore.Channels))
	t.Logf("Members: %d", guildBefore.MemberCount)

	t.Logf("Roles before change: %d", len(guildBefore.Roles))

	_, err = session.GuildRoleCreate(guildBefore.ID)
	if err != nil {
		t.Fatal(err)
	}

	guildAfter, err := session.Guild(mockconstants.TestGuild)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Roles after change: %d", len(guildAfter.Roles))
}

func newState() (*discordgo.State, error) {
	role := mockrole.New(
		mockrole.WithID(mockconstants.TestRole),
		mockrole.WithName(mockconstants.TestRole),
		mockrole.WithPermissions(discordgo.PermissionViewChannel),
	)

	botUser := mockuser.New(
		mockuser.WithID(mockconstants.TestUser+"Bot"),
		mockuser.WithUsername(mockconstants.TestUser+"Bot"),
		mockuser.WithBotFlag(true),
	)

	botMember := mockmember.New(
		mockmember.WithUser(botUser),
		mockmember.WithGuildID(mockconstants.TestGuild),
		mockmember.WithRoles(role),
	)

	userMember := mockmember.New(
		mockmember.WithUser(mockuser.New(
			mockuser.WithID(mockconstants.TestUser),
			mockuser.WithUsername(mockconstants.TestUser),
		)),
		mockmember.WithGuildID(mockconstants.TestGuild),
		mockmember.WithRoles(role),
	)

	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
	)

	privateChannel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestPrivateChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestPrivateChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
		mockchannel.WithPermissionOverwrites(&discordgo.PermissionOverwrite{
			ID:   botMember.User.ID,
			Type: discordgo.PermissionOverwriteTypeMember,
			Deny: discordgo.PermissionViewChannel,
		}),
	)

	return mockstate.New(
		mockstate.WithUser(botUser),
		mockstate.WithGuilds(
			mockguild.New(
				mockguild.WithID(mockconstants.TestGuild),
				mockguild.WithName(mockconstants.TestGuild),
				mockguild.WithRoles(role),
				mockguild.WithChannels(channel, privateChannel),
				mockguild.WithMembers(botMember, userMember),
			),
		),
	)
}
