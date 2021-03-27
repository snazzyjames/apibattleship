package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"NewGame",
		"POST",
		"/games/new",
		NewGame,
	},
	Route{
		"SetupSession",
		"POST",
		"/games/{sessionId}/setup",
		SetupSession,
	},
	Route{
		"PlaySession",
		"POST",
		"/games/{sessionId}/play",
		PlaySession,
	},
	Route{
		"GetSession",
		"GET",
		"/games/{sessionId}",
		GetSession,
	},

	//TODO: Get Game endpoint
}
