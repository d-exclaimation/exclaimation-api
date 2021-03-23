//
//  schema.resolvers.go
//  exclaimation-api
//
//  Created by d-exclaimation on 8:24 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package main

import (
	"github.com/d-exclaimation/exclaimation-api/db"
	"github.com/d-exclaimation/exclaimation-api/graph"
	"github.com/d-exclaimation/exclaimation-api/server"
	"github.com/d-exclaimation/exclaimation-api/server/services"
	"go.uber.org/fx"
)

// Fx Runtime Lifecycle
func main() {
	fx.New(
		fx.Provide(
			// Server application
			server.AppProvider,

			// Postgres Database
			db.EntProvider,

			// Services
			services.PostServiceProvider,

			// GraphQL Module
			graph.ModuleProvider,

			// Handlers / Controllers
			server.AppHandlersProvider,
		),
		fx.Invoke(
			// Gin Middleware and Endpoints Invoker
			server.InvokeMiddleWare,
			server.InvokeHandler,
		),
	).Run()
}
