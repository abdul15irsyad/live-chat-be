package handlers

import (
	"encoding/json"
	"live-chat-be/types"
	"live-chat-be/utils"
	"net/http"
)

type RegisterRequestData struct {
	Name string `json:"name"`
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(writer).Encode(map[string]any{
			"message": "Method Not Allowed",
		})
		return
	}

	var data RegisterRequestData
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	clientNames := utils.MapSlice(utils.Values(clients), func(client types.Client) string {
		return client.Name
	})

	if utils.Includes(clientNames, data.Name) {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{
			"error": "Bad Request",
			"code":  "VALIDATION_ERROR",
			"errors": []map[string]any{
				{
					"field":   "name",
					"message": "Name already exist",
				},
			},
		})
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]any{
		"message": "Register Successfully",
	})
}
