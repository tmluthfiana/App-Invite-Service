package main

import "invitation-app/app"

// @title App Invite API
// @version 1.0
// @description This is an App Invite API Documentation.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.StartApplication()
}
