//
//  graphql.go
//  errors
//
//  Created by d-exclaimation on 5:48 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package errors

import (
	"net/http"
)


func InvalidKeyError() error {
	return NewServiceError(http.StatusForbidden, "Invalid Key")
}

func NotLoggedInError() error {
	return NewServiceError(http.StatusBadRequest, "Not Logged In with Authorized Permission")
}