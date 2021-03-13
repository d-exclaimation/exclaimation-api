//
//  app.go
//  server
//
//  Created by d-exclaimation on 8:05 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package server

import (
	"context"
	"github.com/d-exclaimation/exclaimation-gql/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
	"net/http"
)

const (
	graphqlPath = "/graphql"
	entry = "/"
)

// Fx Provider
func AppProvider(lifecycle fx.Lifecycle) *gin.Engine {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	port := config.GetPort()

	// To Gracefully setup and shuts down http server
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           app,
	}

	// Using Fx Lifecycle create start and stop functions to be invoke at appropriate condition
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				_ = srv.ListenAndServe()
			})()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return app
}

// Fx Invoke Middleware
func InvokeMiddleWare(app *gin.Engine, handlers *AppHandlers) {
	for _, mw := range handlers.Middlewares {
		app.Use(mw)
	}
}

// Fx Invoke Handler
func InvokeHandler(app *gin.Engine, handlers *AppHandlers) {
	app.POST(graphqlPath, handlers.GQLHandler)
	app.GET(entry, handlers.Playground)
}