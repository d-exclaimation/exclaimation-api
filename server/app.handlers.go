//
//  app.handlers.go
//  server
//
//  Created by d-exclaimation on 8:11 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	. "github.com/d-exclaimation/exclaimation-api/server/middleware"
	"github.com/labstack/echo/v4"
)

// AppHandlers / Controller
type AppHandlers struct {
	Middlewares []echo.MiddlewareFunc
	GQLHandler  echo.HandlerFunc
	Playground  echo.HandlerFunc
}

// AppHandlersProvider Fx Provider
func AppHandlersProvider(module generated.Config) *AppHandlers {
	return &AppHandlers{
		Middlewares: []echo.MiddlewareFunc{
			EndPointLoggerMiddleware,
			RateLimiterMiddleware,
			EdgeCaseSecurityMiddleware,
			CorsMiddleware,
			SessionMiddleware,
			EchoContextMiddleware,
		},
		GQLHandler:  GraphqlHandler(module),
		Playground:  PlaygroundHandler(),
	}
}

// GraphqlHandler GraphQL Query Handler
func GraphqlHandler(module generated.Config) echo.HandlerFunc {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(module))
	return func(ctx echo.Context) error {
		graphqlServer.ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PlaygroundHandler Playground Handler
func PlaygroundHandler() echo.HandlerFunc {
	playgroundHandler := playground.Handler("Nodes-Graph API Playground", graphqlPath)
	return func(ctx echo.Context) error {
		playgroundHandler.ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}
