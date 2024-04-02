package mockrest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ewohltman/discordgo-mock/mockconstants"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func (roundTripper *RoundTripper) addHandlersChannels(apiVersion string) {
	pathChannels := fmt.Sprintf("%s/%s", apiVersion, resourceChannels)

	subrouter := roundTripper.router.PathPrefix(pathChannels).Subrouter()

	pathChannelID := "/" + resourceChannelID
	pathChannelIDMessages := fmt.Sprintf("%s/%s", pathChannelID, resourceMessages)
	pathChannelIDInvites := fmt.Sprintf("%s/%s", pathChannelID, resourceInvites)

	getHandlers := subrouter.Methods(http.MethodGet).Subrouter()
	getHandlers.HandleFunc("", roundTripper.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelID, roundTripper.channelsResponseGET)
	getHandlers.HandleFunc(pathChannelIDMessages, roundTripper.channelMessagesResponseGET)

	postHandlers := subrouter.Methods(http.MethodPost).Subrouter()
	postHandlers.HandleFunc(pathChannelIDMessages, roundTripper.channelMessagesResponsePOST)
	postHandlers.HandleFunc(pathChannelIDInvites, roundTripper.channelInvitesResponsePOST)

	deleteHandlers := subrouter.Methods(http.MethodDelete).Subrouter()
	deleteHandlers.HandleFunc(pathChannelID, roundTripper.channelsResponseDelete)

	patchChannels := subrouter.Methods(http.MethodPatch).Subrouter()
	patchChannels.HandleFunc(pathChannelID, roundTripper.channelsResponsePatch)
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

func (roundTripper *RoundTripper) channelsResponseDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)

		return
	}

	sendJSON(w, channel)
}

func (roundTripper *RoundTripper) channelsResponsePatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)

		return
	}

	c := &discordgo.ChannelEdit{}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&c)
	if err != nil {
		sendError(w, err)

		return
	}

	channel.Name = c.Name
	channel.Topic = c.Topic
	channel.MessageCount = *c.Position

	if c.NSFW != nil {
		channel.NSFW = *c.NSFW
	}

	channel.Icon = c.ParentID
	channel.Position = *c.Position
	channel.Bitrate = c.Bitrate
	channel.PermissionOverwrites = c.PermissionOverwrites
	channel.UserLimit = c.UserLimit
	channel.ParentID = c.ParentID

	if c.RateLimitPerUser != nil {
		channel.RateLimitPerUser = *c.RateLimitPerUser
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

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)

		return
	}

	message := &discordgo.Message{}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&message)
	if err != nil {
		sendError(w, err)

		return
	}

	message.ID = randString()

	message.ChannelID = channelID
	channel.LastMessageID = message.ID
	channel.MessageCount++
	channel.Messages = append(channel.Messages, message)

	err = roundTripper.state.MessageAdd(message)
	if err != nil {
		sendError(w, err)

		return
	}

	sendJSON(w, message)
}

func (roundTripper *RoundTripper) channelInvitesResponsePOST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars[resourceChannelIDKey]

	channel, err := roundTripper.state.Channel(channelID)
	if err != nil {
		sendError(w, err)

		return
	}

	guild, err := roundTripper.state.Guild(channel.GuildID)
	if err != nil {
		sendError(w, err)

		return
	}

	invite := &discordgo.Invite{}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&invite)
	if err != nil {
		sendError(w, err)

		return
	}

	invite.Guild = guild
	invite.Channel = channel
	invite.Code = mockconstants.TestInviteCode
	invite.Inviter = roundTripper.state.User

	sendJSON(w, invite)
}
