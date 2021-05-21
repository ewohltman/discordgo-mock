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

type roundTripperFunc func(req *http.Request) (*http.Response, error)

func (rt roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}

type Client struct {
	*http.Client
	state  *discordgo.State
	router *mux.Router
}

func NewClient(state *discordgo.State) *Client {
	router := mux.NewRouter()

	client := &Client{
		Client: &http.Client{
			Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
				responseRecorder := httptest.NewRecorder()
				router.ServeHTTP(responseRecorder, r)
				return responseRecorder.Result(), nil
			}),
		},
		state:  state,
		router: router,
	}

	apiVersion := "/api/v" + discordgo.APIVersion

	client.addHandlersGuilds(apiVersion)
	client.addHandlersChannels(apiVersion)
	client.addHandlersUsers(apiVersion)

	return client
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
