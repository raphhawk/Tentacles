package main

import (
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

	// out, _ := json.MarshalIndent(payload, "", "\t")
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
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
