//
//  schema.resolvers.go
//  exclaimation-gql
//
//  Created by d-exclaimation on 8:24 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package main

import (
	"github.com/d-exclaimation/exclaimation-gql/db"
	"github.com/d-exclaimation/exclaimation-gql/graph"
	"github.com/d-exclaimation/exclaimation-gql/server"
	"github.com/d-exclaimation/exclaimation-gql/server/services"
	"go.uber.org/fx"
)

// Fx Runtime Lifecycle
func main() {
	fx.New(
		fx.Provide(
			// Gin App
			server.AppProvider,

			// Postgres Database
			db.PostgresProvider,

			// Services and Modules
			services.PostServiceProvider,
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
