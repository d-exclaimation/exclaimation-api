//
//  node_factory.go
//  libs
//
//  Created by d-exclaimation on 10:33 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package post

import (
	"fmt"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	"strings"
)

type StringArray []string

func CreateGraphQLNode(val string, i int) *model.PostNode {
	node := "content"
	if strings.HasPrefix(val, "#") {
		node = "header"
	} else if val == "" {
		node = "space"
	}

	return &model.PostNode{
		ID: fmt.Sprintf("PostNode-%d", i),
		Type: node,
		Leaf: val,
	}
}

func ToGraphQLs(t StringArray) []*model.PostNode {
	res := make([]*model.PostNode, len(t))
	for i, val := range t {
		res[i] = CreateGraphQLNode(val, i)
	}
	return res
}
