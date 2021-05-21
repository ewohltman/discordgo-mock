package mockrest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func (client *Client) addHandlersUsers(apiVersion string) {
	pathUsers := fmt.Sprintf("%s/%s", apiVersion, resourceUsers)

	subrouter := client.router.PathPrefix(pathUsers).Subrouter()

	pathUserID := fmt.Sprintf("/%s", resourceUserID)

	getHandlers := subrouter.Methods(http.MethodGet).Subrouter()
	getHandlers.HandleFunc("", client.usersResponse)
	getHandlers.HandleFunc(pathUserID, client.usersResponse)
}

func (client *Client) usersResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars[resourceUserIDKey]

	user := &discordgo.User{
		ID: userID,
	}

	respBody, err := json.Marshal(user)
	if err != nil {
		sendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respBody)
}
