//
//  raw_external.go
//  ent
//
//  Created by d-exclaimation on 12:52 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package ent

import (
	"database/sql"
	esql "entgo.io/ent/dialect/sql"
)

func (c *Client) DB() *sql.DB {
	return c.driver.(*esql.Driver).DB()
}
