//
//  rate_limiter.go
//  middleware
//
//  Created by d-exclaimation on 8:54 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return em.RateLimiter(em.NewRateLimiterMemoryStore(20))(next)
}
