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

	//	New Game
	//	Method: POST
	//	Request:
	//		JSON {"player_one": "<name>", "player_two": "<other name>"}
	//	Response:
	//		JSON {"session_id: "<unique session id>", "phase":"setup", "player":"<player_one_name>"}
	Route{
		"NewGame",
		"POST",
		"/games/new",
		NewGame,
	},
	//	Setup Game
	//	Method: POST
	//	Path param: sessionId (the ID of the game to setup)
	//	Request:
	//		JSON {"ship": "carrier|battleship|cruiser|submarine|destroyer", "coordinate": "<A0-G9>",
	//		direction:"down|right", player: "<player_name>"}
	//	Response:
	//		if still in setup: JSON {"placed: "true|false", "next_player": "<next_player>", "phase":"setup"}
	//		if ready for play: JSON {"placed: "true|false", "phase":"play"}
	Route{
		"SetupGame",
		"POST",
		"/games/{sessionId}/setup",
		SetupGame,
	},
	//	Play Game
	//	Method: POST
	//	Path Param: sessionId (the ID of the game to play)
	//	Request:
	//		JSON {"coordinate":"<A0-G9>", "player": "<player_name"}
	//	Response:
	//		JSON {"result":"hit|miss|hit_sunk|hit_good_game|not_your_turn|game_over",
	//		"next_player": "<next_player>"}
	Route{
		"PlayGame",
		"POST",
		"/games/{sessionId}/play",
		PlayGame,
	},
	// Get Game
	// Method: GET
	// Path Param: sessionId (the ID of the game to get)
	// Response:
	// 		JSON {"phase": "setup|play|game_over", "players":["<player_one>","<player_two>"]
	Route{
		"GetGame",
		"GET",
		"/games/{sessionId}",
		GetGame,
	},
}
