package mockrest

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

const (
	resourceUsers     = "users"
	resourceUserIDKey = "userID"
	resourceUserID    = "{" + resourceUserIDKey + "}"

	resourceRoles     = "roles"
	resourceRoleIDKey = "roleID"
	resourceRoleID    = "{" + resourceRoleIDKey + "}"

	resourceChannels     = "channels"
	resourceChannelIDKey = "channelID"
	resourceChannelID    = "{" + resourceChannelIDKey + "}"
	resourceMessages     = "messages"

	resourceGuilds     = "guilds"
	resourceGuildIDKey = "guildID"
	resourceGuildID    = "{" + resourceGuildIDKey + "}"

	resourceMembers = "members"
)

type RoundTripper struct {
	router *mux.Router
	state  *discordgo.State
}

func (roundTripper *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	responseRecorder := httptest.NewRecorder()
	roundTripper.router.ServeHTTP(responseRecorder, req)
	return responseRecorder.Result(), nil
}

func NewTransport(state *discordgo.State) http.RoundTripper {
	router := mux.NewRouter()

	roundTripper := &RoundTripper{
		router: router,
		state:  state,
	}

	apiVersion := "/api/v" + discordgo.APIVersion

	roundTripper.addHandlersGuilds(apiVersion)
	roundTripper.addHandlersChannels(apiVersion)
	roundTripper.addHandlersUsers(apiVersion)

	return roundTripper
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}

func randString() string {
	const (
		letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		stringLength = 10
	)

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, stringLength)

	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
