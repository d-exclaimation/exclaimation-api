//
//  connection.go
//  db
//
//  Created by d-exclaimation on 7:18 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package db

import (
	"context"
	"database/sql"
	"github.com/d-exclaimation/exclaimation-gql/config"
	"github.com/d-exclaimation/exclaimation-gql/db/entities"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gopkg.in/gorp.v1"
	"log"
)

// Fx Provider
func PostgresProvider(lifecycle fx.Lifecycle) *gorp.DbMap {
	// Open SQL Connection
	db, err := sql.Open("postgres", config.GetDatabaseURL())
	if err != nil {
		log.Fatalln(err)
	}

	// Create Relational Persistence
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbMap.AddTableWithName(entities.Post{}, "posts").SetKeys(true, "Id")
	dbMap.TraceOff()

	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln(err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return dbMap.Db.Close()
		},
	})

	return dbMap
}
