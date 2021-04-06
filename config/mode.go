//
//  mode.go
//  config
//
//  Created by d-exclaimation on 9:37 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import "os"

type EnvMode string

const (
	Prod        EnvMode = "prod"
	Dev         EnvMode = "dev"
	Maintenance EnvMode = "down"
)

func GetServerMode() EnvMode {
	mode := os.Getenv("ENV_MODE")
	switch mode {
	case "PRODUCTION":
		return Prod
	case "DEVELOPMENT":
		return Dev
	case "MAINTENANCE":
		return Maintenance
	default:
		return Dev
	}
}
