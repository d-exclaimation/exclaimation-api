//
//  artifacts.go
//  config
//
//  Created by d-exclaimation on 11:03 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import "os"

func GetKey() string {
	key := os.Getenv("ACCESS_TOKEN")
	if key == "" || len(key) < 1 {
		key = "no-key"
	}
	return key
}

func GetSessionSecret() string {
	secret, ok := os.LookupEnv("SESSION_SECRET")
	if !ok {
		secret = "NO"
	}
	return secret
}


func GetComputedKey() string {
	comp, ok := os.LookupEnv(GetKey())
	if !ok {
		comp = "NO"
	}
	return comp
}