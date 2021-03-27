//
//  logger.go
//  middleware
//
//  Created by d-exclaimation on 1:07 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

func loggerFormat() string {
	var (
		mode = config.GetServerMode()
		header = ""
		method = ""
		hexColor = ""
		endpoint = ""
	)
	switch mode {
	case config.Prod:
		hexColor = "#fc038c"
		method = " ${status} ${method} "
		endpoint = " >> ${uri}\n"
		break
	case config.Maintenance:
		header = "${time_rfc3339_nano} |"
		hexColor = "#6703fc"
		method = " ${status} "
		endpoint = "| ${latency_human} |\n"
		break
	default:
		header = "${time_rfc3339_nano} |"
		hexColor = "#20bcaf"
		method = " ${status} ${method} "
		endpoint = "| ${latency_human} | >> ${uri}\n"
		break
	}

	return header +
		color.NewRGBStyle(
			color.RGB(200, 200, 200),
			color.HEX(hexColor, true),
		).
		Sprint(method) +
		endpoint
}

func EndPointLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return em.LoggerWithConfig(em.LoggerConfig{
		Format: loggerFormat(),
	})(next)
}
