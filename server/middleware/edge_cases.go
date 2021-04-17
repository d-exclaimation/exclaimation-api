//
//  edge_cases.go
//  middleware
//
//  Created by d-exclaimation on 8:58 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

func EdgeCaseSecurityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// TODO: Create proper config
	return em.Secure()(next)
}
