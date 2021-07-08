package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/repository"

	"github.com/lutianzhou001/neo3fura_graphql/graph/generated"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
)

var (
	blockRepository repository.BlockRepository = repository.NewBlockRepository()
)


func (r *queryResolver) Blocks(ctx context.Context) ([]*model.Block, error) {
	return blockRepository.FindAll(),nil;
}

func (r *queryResolver) LastBlock(ctx context.Context) (*model.Block, error) {
	return blockRepository.FindLastBlock(),nil;
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

