package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/d-exclaimation/exclaimation-api/utils/pipes"
	"github.com/d-exclaimation/exclaimation-api/utils/slice"
	"strings"

	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
)

func (r *postResolver) Snippet(_ context.Context, obj *model.Post) (string, error) {
	snippet := strings.Split(obj.Body, "")
	if len(snippet) > 60 {
		snippet = snippet[0:60]
	}
	return strings.Join(snippet, ""), nil
}

func (r *postResolver) Nodes(_ context.Context, obj *model.Post) ([]*model.PostNode, error) {
	leaves := slice.ReduceStr(strings.Split(obj.Body, "\n"), pipes.IsolateReducer([]pipes.IsolateStringPipe{
		func(row string) bool { return strings.HasPrefix(row, "#") },
		func(row string) bool { return row == "" },
	}))
	res := make([]*model.PostNode, len(leaves))
	for i, val := range leaves {
		node := "content"
		if strings.HasPrefix(val, "#") {
			node = "header"
		} else if val == "" {
			node = "space"
		}
		
		res[i] = &model.PostNode{
			Type: node,
			Leaf: val,
		}
	}
	return res, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
