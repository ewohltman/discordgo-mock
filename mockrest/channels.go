package mockrest

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func (roundTripper *RoundTripper) addHandlersChannels(apiVersion string) {
	pathChannels := fmt.Sprintf("%s/%s", apiVersion, resourceChannels)

	subrouter := roundTripper.router.PathPrefix(pathChannels).Subrouter()

	pathChannelID := fmt.Sprintf("/%s", resourceChannelID)
	pathChannelIDMessages := fmt.Sprintf("%s/%s", pathChannelID, resourceMessages)

	getHandlers := subrouter.Methods(http.MethodGet).Subrouter()
	getHandlers.HandleFunc("", roundTripper.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelID, roundTripper.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelIDMessages, roundTripper.channelMessagesResponseGET)

	postHandlers := subrouter.Methods(http.MethodPost).Subrouter()
	postHandlers.HandleFunc(pathChannelIDMessages, roundTripper.channelMessagesResponsePOST)
}

func (roundTripper *RoundTripper) channelsResponseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)
		return
	}

	sendJSON(w, channel)
}

func (roundTripper *RoundTripper) channelMessagesResponseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)
		return
	}

	sendJSON(w, channel.Messages)
}

func (roundTripper *RoundTripper) channelMessagesResponsePOST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	message := &discordgo.Message{
		ID:        randString(),
		ChannelID: channelID,
	}

	sendJSON(w, message)
}
