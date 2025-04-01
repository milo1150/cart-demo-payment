package main

import (
	"github.com/labstack/echo/v4"
	"github.com/milo1150/cart-demo-payment/internal/database"
	"github.com/milo1150/cart-demo-payment/internal/grpc"
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

	// NATS JetStream
	js := nats.ConnectJetStream(nc)

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Initialize Zap Logger
	logger := middlewares.InitializeZapLogger()

	// Global state
	appState := &types.AppState{
		DB:   db,
		NATS: nc,
		Log:  logger,
		JS:   js,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Run NATS services
	go nats.SubscribeCheckoutEvent(js, logger, db)

	// gRPC Servers
	go grpc.StartPaymentGRPCServer(appState)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
