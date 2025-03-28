package main

import (
	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/database"
	"github.com/milo1150/cart-demo-payment/internal/loader"
	"github.com/milo1150/cart-demo-payment/internal/middlewares"
	"github.com/milo1150/cart-demo-payment/internal/nats"
	"github.com/milo1150/cart-demo-payment/internal/routes"
	"github.com/milo1150/cart-demo-payment/internal/types"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// NATS
	nc := nats.ConnectNATS()
	defer nc.Close()

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Initialize Zap Logger
	logger := middlewares.InitializeZapLogger()

	// Connect to gRPC Servers

	// Global state
	appState := &types.AppState{
		DB:   db,
		NATS: nc,
		Log:  logger,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Init NATS Pub/Sub
	nats.SubscribeToUserService(nc, logger, db)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
