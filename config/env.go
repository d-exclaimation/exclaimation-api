//
//  env.go
//  config
//
//  Created by d-exclaimation on 7:19 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InvokeDotEnv() {
	if GetServerMode() == Prod {
		return
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}
}

const defaultPort = "4000"

func GetDatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" || len(dbURL) < 1 {
		dbURL = "postgres://127.0.0.1:5432/dev-site?sslmode=disable"
	}
	return dbURL
}


func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" || len(port) < 1 {
		port = defaultPort
	}
	return port
}

func GetKey() string {
	key := os.Getenv("ACCESS_TOKEN")
	if key == "" || len(key) < 1 {
		key = "no-key"
	}
	return key
}
