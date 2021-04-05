package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/graph/libs"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	"github.com/d-exclaimation/exclaimation-api/utils/pipes"
	"github.com/d-exclaimation/exclaimation-api/utils/slice"
)

func (r *postResolver) Snippet(ctx context.Context, obj *model.Post) (string, error) {
	snippet := strings.Split(obj.Body, "")
	if len(snippet) > 60 {
		snippet = snippet[0:60]
	}
	return strings.Join(snippet, ""), nil
}

func (r *postResolver) Nodes(ctx context.Context, obj *model.Post) ([]*model.PostNode, error) {
	leaves := libs.StringArray(
		slice.ReduceStr(
			strings.Split(obj.Body, "\n"),
			pipes.IsolateReducer(
				[]pipes.IsolateStringPipe{
					func(row string) bool { return strings.HasPrefix(row, "#") },
					func(row string) bool { return row == "" },
				},
			),
		)).ToGraphQLs()
	return leaves, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
