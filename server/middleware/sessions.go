//
//  sessions.go
//  middleware
//
//  Created by d-exclaimation on 10:06 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/gorilla/sessions"
	ss "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return ss.Middleware(sessions.NewCookieStore([]byte(config.GetSessionSecret())))(next)
}
