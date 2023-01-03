package core

import "travel_planning_assistant/initialize"

type server struct {
}

func RunWindowsServer() {
	router := initialize.Routers()
	router.Static("/from-generator", "./resource/page")
}
