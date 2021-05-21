package mockrest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func (client *Client) addHandlersChannels(apiVersion string) {
	pathChannels := fmt.Sprintf("%s/%s", apiVersion, resourceChannels)

	subrouter := client.router.PathPrefix(pathChannels).Subrouter()

	pathChannelID := fmt.Sprintf("/%s", resourceChannelID)
	pathChannelIDMessages := fmt.Sprintf("%s/%s", pathChannelID, resourceMessages)

	getHandlers := subrouter.Methods(http.MethodGet).Subrouter()
	getHandlers.HandleFunc("", client.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelID, client.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelIDMessages, client.channelMessagesResponseGET)

	postHandlers := subrouter.Methods(http.MethodPost).Subrouter()
	postHandlers.HandleFunc(pathChannelIDMessages, client.channelMessagesResponsePOST)
}

func (client *Client) channelsResponseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := client.state.Channel(channelID)
	if err != nil {
		sendError(w, err)
		return
	}

	respBody, err := json.Marshal(channel)
	if err != nil {
		sendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}

func (client *Client) channelMessagesResponseGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := client.state.Channel(channelID)
	if err != nil {
		sendError(w, err)
		return
	}

	respBody, err := json.Marshal(channel.Messages)
	if err != nil {
		sendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}

func (client *Client) channelMessagesResponsePOST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	message := &discordgo.Message{
		ID:        randString(),
		ChannelID: channelID,
	}

	respBody, err := json.Marshal(message)
	if err != nil {
		sendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}
