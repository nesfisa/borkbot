package borkbot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/form"
)

func decodeFetchBorkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var br fetchBorkRequest
	r.ParseForm()
	decoder := form.NewDecoder()
	err := decoder.Decode(&br, r.Form)
	if err != nil {
		return nil, err
	}
	return br, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// SlackForm contains all of the information that comes in from a slack post request
// Requests from slack come as application/x-www-form-urlencoded so we need a way to
// decode these values when receving the request
type SlackForm struct {
	UserID      string `form:"user_id"`
	UserName    string `form:"user_name"`
	TeamDomain  string `form:"team_domain"`
	ResponseURL string `form:"response_url"`
	TriggerID   string `form:"trigger_id"`
	Text        string `form:"text"`
	Token       string `form:"token"`
	ChannelName string `form:"channel_name"`
	Command     string `form:"command"`
}
