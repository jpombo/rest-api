package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"rest-api/internal/models"
)

func main() {

	// create data to send
	request := models.CreateUserRequest{
		Name:  "Person3",
		Email: "person3@person.com",
	}

	// parser data to json
	requestParsed, errParse := json.Marshal(request)
	slog.Info("antes parser")
	if errParse == nil {
		slog.Info("antes post")
		responseData, errPost := http.Post(
			"http://127.0.0.1:8080/addusers",
			"application/json",
			bytes.NewReader(requestParsed),
		)
		slog.Info("antes retorno post")
		if errPost == nil && responseData.StatusCode == http.StatusCreated {
			var dataResponse models.CreateUserResponse
			errResultParsed := json.NewDecoder(responseData.Body).Decode(&dataResponse)
			if errResultParsed == nil {
				fmt.Println("User created with id: ", dataResponse)
			} else {
				slog.Error("Error in data parser to receive", "error", errResultParsed.Error())
			}
		} else {
			if responseData.StatusCode == http.StatusBadRequest {
				var badRequestResponse models.ErrorResponse
				errResultParsed := json.NewDecoder(responseData.Body).Decode(&badRequestResponse)
				if errResultParsed == nil {
					fmt.Println(badRequestResponse)
				} else {
					slog.Error("Error in data parser to receive", "error", errResultParsed.Error())
				}
			} else {
				slog.Error("Error in data  parser to send", "error", errPost.Error())
			}
		}
	} else {
		slog.Error("Error in data  parser to send", "error", errParse.Error())
	}
}
