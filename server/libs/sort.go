//
//  sort.go
//  libs
//
//  Created by d-exclaimation on 1:45 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package libs

import (
	"github.com/d-exclaimation/exclaimation-api/ent"
	"github.com/d-exclaimation/exclaimation-api/ent/post"
)

type SortBy string

const (
	ByLatest SortBy = "byLatest"
	ByHot    SortBy = "byHot"
)

func ToSortBy(sorting string) SortBy {
	switch sorting {
	case "latest":
		return ByLatest
	case "hot":
		return ByHot
	case "recent":
		return ByLatest
	case "top":
		return ByHot
	default:
		return ByLatest
	}
}

func EntOrderBy(sorting SortBy) ent.OrderFunc {
	switch sorting {
	case ByLatest:
		return ent.Desc(post.FieldID)
	case ByHot:
		return ent.Desc(post.FieldCrabrave)
	default:
		return ent.Desc(post.FieldID)
	}
}
