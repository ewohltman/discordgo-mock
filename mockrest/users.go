package mockrest

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func (roundTripper *RoundTripper) addHandlersUsers(apiVersion string) {
	pathUsers := fmt.Sprintf("%s/%s", apiVersion, resourceUsers)

	subrouter := roundTripper.router.PathPrefix(pathUsers).Subrouter()

	pathUserID := "/" + resourceUserID
	pathUserIDChannels := fmt.Sprintf("/%s/channels", resourceUserID)

	getHandlers := subrouter.Methods(http.MethodGet).Subrouter()
	getHandlers.HandleFunc("", roundTripper.usersResponse)
	getHandlers.HandleFunc(pathUserID, roundTripper.usersResponse)

	postHandlers := subrouter.Methods(http.MethodPost).Subrouter()
	postHandlers.HandleFunc(pathUserIDChannels, roundTripper.userChannelsResponse)
}

func (roundTripper *RoundTripper) usersResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars[resourceUserIDKey]

	sendJSON(w, &discordgo.User{ID: userID})
}

func (roundTripper *RoundTripper) userChannelsResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars[resourceUserIDKey]

	ch := &discordgo.Channel{ID: userID,
		Type: discordgo.ChannelTypeDM,
	}

	err := roundTripper.state.ChannelAdd(ch)
	if err != nil {
		sendError(w, err)

		return
	}

	sendJSON(w, &discordgo.Channel{ID: userID})
}
