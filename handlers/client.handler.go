package handlers

import (
	"encoding/json"
	"live-chat-be/types"
	"live-chat-be/utils"
	"net/http"
)

func GetAllClientHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(writer).Encode(map[string]any{
			"message": "Method Not Allowed",
		})
		return
	}

	clientNames := utils.MapSlice(utils.Values(clients), func(client types.Client) string {
		return client.Name
	})

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]any{
		"message": "Get All Client Successfully",
		"data": map[string]any{
			"clientNames": clientNames,
		},
	})
}
