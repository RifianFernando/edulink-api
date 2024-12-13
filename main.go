// package handler
package main

import (
	"github.com/edulink-api/config"
	"github.com/edulink-api/connections"
	"github.com/edulink-api/lib"
	_ "github.com/edulink-api/request"
	"github.com/edulink-api/routes"
)

// init function for environment setup
func init() {
	connections.LoadEnvVariables()
	config.InitializeSessionStore()

	// Initialize database connection
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

// Vercel requires an HTTP handler function that serves the API
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	app := routes.SetupRouter()

// 	app.ServeHTTP(w, r)
// }

func main() {
	// Set up the router
	r := routes.SetupRouter()

	// Run the server
	err := r.Run() // listen and serve on
	if err != nil {
		lib.HandleError(err, "Failed to start server")
	}
}
