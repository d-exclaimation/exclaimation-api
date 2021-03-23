//
//  graphql.go
//  errors
//
//  Created by d-exclaimation on 5:48 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package errors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

func (err *ServiceError) ToGQLError() error {
	return gqlerror.Errorf("(%d) %s", err.Type, err.Response)
}

func InvalidKeyError() error {
	return NewServiceError(http.StatusForbidden, "Invalid Key").ToGQLError()
}