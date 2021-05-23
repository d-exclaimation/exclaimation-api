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
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

const (
	graphqlPath = "/graphql"
	entry = "/"
)

// AppProvider Fx Provider
func AppProvider(lifecycle fx.Lifecycle) *echo.Echo {
	app := echo.New()
	port := config.GetPort()

	// Using Fx Lifecycle create start and stop functions to be invoke at appropriate condition
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				_ = app.Start(":" + port)
			})()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown(ctx)
		},
	})

	return app
}

// InvokeMiddleWare Fx Invoke Middleware
func InvokeMiddleWare(app *echo.Echo, handlers *AppHandlers) {
	for _, mw := range handlers.Middlewares {
		app.Use(mw)
	}
}

// InvokeHandler Fx Invoke Handler
func InvokeHandler(app *echo.Echo, handlers *AppHandlers) {
	// Assign the uri, to the graphql handler
	app.POST(graphqlPath, handlers.GQLHandler)

	// Assign playground only for non-prod, only (some people doesn't spam on the browser)
	if config.GetServerMode() == config.Prod {
		app.GET(entry, redirectConnection)
		app.GET("/index.php", redirectConnection)
		app.GET("/.env", redirectConnection)
		app.GET("/console/", redirectConnection)
		app.GET("/wp-admin", redirectConnection)
	} else {
		app.GET(entry, handlers.Playground)
	}
}

func redirectConnection(ctx echo.Context) error {
	return ctx.Redirect(http.StatusPermanentRedirect, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}
