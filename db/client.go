//
//  client.go
//  db
//
//  Created by d-exclaimation on 9:46 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package db

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/d-exclaimation/exclaimation-api/ent"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
)

func EntProvider(lifecycle fx.Lifecycle) *ent.Client {
	client, err := ent.Open(dialect.Postgres, config.GetDatabaseURL())
	if err != nil {
		log.Fatalln(err.Error())
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
		OnStop: func(_ context.Context) error {
			return client.Close()
		},
	})

	return client
}

