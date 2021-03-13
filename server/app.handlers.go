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
	"github.com/d-exclaimation/exclaimation-gql/graph/generated"
	"github.com/d-exclaimation/exclaimation-gql/server/middleware"
	"github.com/gin-gonic/gin"
)

// AppHandlers / Controller
type AppHandlers struct {
	Middlewares []gin.HandlerFunc
	GQLHandler  gin.HandlerFunc
	Playground  gin.HandlerFunc
}

// Fx Provider
func AppHandlersProvider(module generated.Config) *AppHandlers {
	return &AppHandlers{
		Middlewares: []gin.HandlerFunc{middleware.GinContextToContextMiddleware()},
		GQLHandler:  GraphqlHandler(module),
		Playground:  PlaygroundHandler(),
	}
}

// GraphQL Query Handler
func GraphqlHandler(module generated.Config) gin.HandlerFunc {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(module))
	return func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// Playground Handler
func PlaygroundHandler() gin.HandlerFunc {
	playgroundHandler := playground.Handler("Nodes-Graph API Playground", graphqlPath)
	return func(ctx *gin.Context) {
		playgroundHandler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
