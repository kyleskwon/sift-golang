// Copyright 2016 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// MIT License

package sift

import "net/http"

// Label
type Label struct {
	IsBad       bool     `json:"is_bad,omitempty"`
	Time        int      `json:"time,omitempty"`
	Reasons     []string `json:"reasons,omitempty"`
	Description string   `json:"description,omitempty"`
}

// Reason
type Reason struct {
	Name    string `json:"name,omitempty"`
	Value   int    `json:"value,omitempty"`
	Details struct {
		Users string `json:"users,omitempty"` // TODO: are there other possible keys besides "users"?
	} `json:"details,omitempty"`
}

// Trigger
type Trigger struct {
	Type    string `json:"type,omitempty"`
	Source  string `json:"source,omitempty"`
	Trigger struct {
		ID string `json:"id,omitempty"`
	} `json:"trigger,omitempty"`
}

// Action
type Action struct {
	ID     string `json:"id,omitempty"`
	Action struct {
		ID string `json:"id,omitempty"`
	} `json:"action,omitempty"`
	Entity struct {
		ID string `json:"id,omitempty"`
	}
	Time     int       `json:"time,omitempty"`
	Triggers []Trigger `json:"triggers,omitempty"`
}

// Response
type Response struct {
	// Used for debugging purposes I guess....
	HTTPStatus       string      `json:"-"`
	HTTPStatusCode   int         `json:"-"`
	HTTPStatusHeader http.Header `json:"-"`
	HTTPResponseBody string      `json:"-"`
	// ------------------------------------------------------------

	Status        int            `json:"status,omitempty"`
	ErrorMessage  string         `json:"error_message,omitempty"`
	Request       string         `json:"request,omitempty"`
	Time          int            `json:"time,omitempty"`
	ScoreResponse *ScoreResponse `json:"score_response,omitempty"`
}

// ScoreResponse
type ScoreResponse struct {
	Status       int      `json:"status,omitempty"`
	ErrorMessage string   `json:"error_message,omitempty"`
	UserID       string   `json:"user_id,omitempty"`
	LatestLabel  *Label   `json:"latest_label,omitempty"`
	Score        float64  `json:"score,omitempty"`
	Reasons      []Reason `json:"reasons,omitempty"`
	Actions      []Action `json:"actions,omitempty"`
}
