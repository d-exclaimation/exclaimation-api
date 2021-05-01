//
//  sessions_context.go
//  libs
//
//  Created by d-exclaimation on 10:52 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package libs

import (
	"github.com/d-exclaimation/exclaimation-api/config"
	e "github.com/d-exclaimation/exclaimation-api/server/errors"
	"github.com/d-exclaimation/exclaimation-api/server/middleware"
	"github.com/gorilla/sessions"
	ss "github.com/labstack/echo-contrib/session"
	"golang.org/x/net/context"
)

func fetchSessions(ctx context.Context) (*sessions.Session, error) {
	c, err := middleware.EchoFromContext(ctx)
	if err != nil {
		return nil, err
	}
	sess, err := ss.Get(config.GetSessionSecret(), c)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// GiveTheCookie send back cookie for sessions auth
func GiveTheCookie(ctx context.Context, val string) (string, error) {
	c, err := middleware.EchoFromContext(ctx)
	if err != nil {
		return "", err
	}
	sess, err := ss.Get(config.GetSessionSecret(), c)
	if err != nil {
		return "", err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24 * 365,
		HttpOnly: true,
	}
	sess.Values[val] = config.GetComputedKey()
	err = sess.Save(c.Request(), c.Response())
	return "OK", err
}

// StealTheCookie grab the session token from client cookie jar
func StealTheCookie(ctx context.Context, val string) (*string, error) {
	sess, err := fetchSessions(ctx)
	if err != nil {
		return nil, err
	}
	if stamp, ok := sess.Values[val]; ok {
		res := stamp.(string)
		return &res, nil
	}
	return nil, e.NotLoggedInError()
}