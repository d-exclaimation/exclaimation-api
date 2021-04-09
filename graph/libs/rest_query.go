//
//  rest_query.go
//  libs
//
//  Created by d-exclaimation on 4:43 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package libs

import (
	"encoding/json"
	"net/http"
)

func GetAndParse(url string, parsed interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer (func() {
		_ = resp.Body.Close()
	})()

	return json.NewDecoder(resp.Body).Decode(parsed)
}
