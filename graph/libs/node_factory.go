//
//  node_factory.go
//  libs
//
//  Created by d-exclaimation on 10:33 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package libs

import (
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	"strings"
)

type StringArray []string

func CreateGraphQLNode(val string) *model.PostNode {
	node := "content"
	if strings.HasPrefix(val, "#") {
		node = "header"
	} else if val == "" {
		node = "space"
	}

	return &model.PostNode{
		Type: node,
		Leaf: val,
	}
}

func (t StringArray) ToGraphQLs() []*model.PostNode {
	res := make([]*model.PostNode, len(t))
	for i, val := range t {
		res[i] = CreateGraphQLNode(val)
	}
	return res
}
