package mocksession_test

import (
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

type testFunc func(t *testing.T, session *discordgo.Session, expectedResources *resources)

type resources struct {
	user    *discordgo.User
	role    *discordgo.Role
	member  *discordgo.Member
	channel *discordgo.Channel
	guild   *discordgo.Guild
}

func TestNew(t *testing.T) {
	expectedResources := setupResources()
	session := newSession(expectedResources)

	type tc struct {
		name string
		run  testFunc
	}

	testCases := []*tc{
		{name: "user", run: testUser},
		{name: "channel", run: testChannel},
		{name: "guild", run: testGuild},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			testCase.run(t, session, expectedResources)
		})
	}
}

func setupResources() *resources {
	user := mockuser.New(
		mockuser.WithID(mockconstants.TestUser),
		mockuser.WithUsername(mockconstants.TestUser),
	)

	role := mockrole.New(
		mockrole.WithID(mockconstants.TestRole),
		mockrole.WithName(mockconstants.TestRole),
		mockrole.WithPermissions(discordgo.PermissionViewChannel),
	)

	member := mockmember.New(
		mockmember.WithUser(user),
		mockmember.WithGuildID(mockconstants.TestGuild),
		mockmember.WithRoles(role),
	)

	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
		mockchannel.WithPermissionOverwrites(&discordgo.PermissionOverwrite{
			ID:   member.User.ID,
			Type: discordgo.PermissionOverwriteTypeMember,
			Deny: discordgo.PermissionViewChannel,
		}),
	)

	guild := mockguild.New(
		mockguild.WithID(mockconstants.TestGuild),
		mockguild.WithName(mockconstants.TestGuild),
		mockguild.WithRoles(role),
		mockguild.WithChannels(channel),
		mockguild.WithMembers(member),
	)

	return &resources{
		user:    user,
		role:    role,
		member:  member,
		channel: channel,
		guild:   guild,
	}
}

func newSession(expectedResources *resources) *discordgo.Session {
	state, err := mockstate.New(
		mockstate.WithUser(mockuser.New(
			mockuser.WithID("Bot"),
			mockuser.WithUsername("Bot"),
			mockuser.IsBot(true),
		)),
		mockstate.WithGuilds(expectedResources.guild),
	)
	if err != nil {
		panic(err)
	}

	session, err := mocksession.New(
		mocksession.WithState(state),
		mocksession.WithRESTClient(mockrest.NewClient(state)),
	)
	if err != nil {
		panic(err)
	}

	return session
}

func testUser(t *testing.T, session *discordgo.Session, expectedResources *resources) {
	foundUser, err := session.User(expectedResources.user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if foundUser.ID != expectedResources.user.ID {
		t.Fatalf("unexpected user id: %s", foundUser.ID)
	}
}

func testChannel(t *testing.T, session *discordgo.Session, expectedResources *resources) {
	foundChannel, err := session.Channel(expectedResources.channel.ID)
	if err != nil {
		t.Fatal(err)
	}

	if foundChannel.ID != expectedResources.channel.ID {
		t.Fatalf("unexpected channel id: %s", foundChannel.ID)
	}
}

func testGuild(t *testing.T, session *discordgo.Session, expectedResources *resources) {
	foundGuild, err := session.Guild(expectedResources.guild.ID)
	if err != nil {
		t.Fatal(err)
	}

	if foundGuild.ID != expectedResources.guild.ID {
		t.Fatalf("unexpected guild id: %s", foundGuild.ID)
	}

	foundMember, err := session.GuildMember(expectedResources.guild.ID, expectedResources.user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if foundMember.User.ID != expectedResources.member.User.ID {
		t.Errorf("unexpected member user id: %s", foundMember.User.ID)
	}

	var roleWasFound bool

	foundRoles, err := session.GuildRoles(expectedResources.guild.ID)
	if err != nil {
		t.Fatal(err)
	}

	for _, foundRole := range foundRoles {
		if foundRole.ID == expectedResources.role.ID {
			roleWasFound = true
			break
		}
	}

	if !roleWasFound {
		t.Fatalf("role not found in guild: %s", expectedResources.role.ID)
	}

	createdRole, err := session.GuildRoleCreate(expectedResources.guild.ID)
	if err != nil {
		t.Fatal(err)
	}

	const expectedColor = 1

	createdRole, err = session.GuildRoleEdit(
		expectedResources.guild.ID,
		createdRole.ID,
		mockconstants.TestRole,
		expectedColor,
		false,
		0,
		false,
	)
	if err != nil {
		t.Fatal(err)
	}

	if createdRole.Color != expectedColor {
		t.Fatalf("unexpected role color: %d", createdRole.Color)
	}
}
