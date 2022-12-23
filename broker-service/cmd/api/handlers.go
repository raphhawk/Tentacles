package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (app *Config) Broker(
	w http.ResponseWriter,
	r *http.Request,
) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

func (app *Config) authenticate(
	w http.ResponseWriter,
	a AuthPayload,
) {
	// create json data to be send to auth service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
  request, err := http.NewRequest(
    "POST",
    "http://authentication-service/authenticate",
    bytes.NewBuffer(jsonData),
  )
  if err != nil {
    app.errorJSON(w, err)
    return
  }

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    app.errorJSON(w, err)
    return
  }
  defer response.Body.Close()

	// make sure we get back the correct status code
  if response.StatusCode == http.StatusUnauthorized {
    app.errorJSON(w, errors.New("invalid credentials"))
    return
  } else if response.StatusCode != http.StatusAccepted {
    app.errorJSON(w, errors.New("error calling auth service"))
    return
  }

  var jsonFromService jsonResponse
  err = json.NewDecoder(response.Body).Decode(&jsonFromService)
  if err != nil {
    app.errorJSON(w, err)
    return
  }

  if jsonFromService.Error {
    app.errorJSON(w, err, http.StatusUnauthorized)
    return
  }

  var payload jsonResponse 
  payload.Error = false
  payload.Message = "Authenticated!"
  payload.Data = jsonFromService.Data 

  app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) HandleSubmission(
	w http.ResponseWriter,
	r *http.Request,
) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(err)
		return
	}

	switch requestPayload.Action {
	case "auth":
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}
