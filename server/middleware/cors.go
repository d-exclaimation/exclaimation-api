//
//  cors.go
//  middleware
//
//  Created by d-exclaimation on 1:10 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func allowedMethods() []string {
	switch config.GetServerMode() {
	case config.Prod:
		return []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodConnect}
	default:
		return []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete}
	}
}

func CorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return em.CORSWithConfig(em.CORSConfig{
		Skipper:      em.DefaultSkipper,
		AllowCredentials: config.GetServerMode() != config.Maintenance,
		AllowOrigins: []string{ "http://localhost:3000" },
		AllowMethods: allowedMethods(),
	})(next)
}
