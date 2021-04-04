//
//  post.go
//  model
//
//  Created by d-exclaimation on 2:52 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Crabrave int    `json:"crabrave"`
}
