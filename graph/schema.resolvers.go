package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lutianzhou001/neo3fura_graphql/graph/generated"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"github.com/lutianzhou001/neo3fura_graphql/repository"
)

func (r *queryResolver) BlockByBlockHash(ctx context.Context, blockhash *string) (*model.Block, error) {
	return blockRepository.BlockByBlockHash(blockhash)
}

func (r *queryResolver) Nep11TransferNotificationsByAddress(ctx context.Context, address *string) ([]*model.Nep11TransferNotification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nep11TransferNotificationByTransactionHash(ctx context.Context, transactionhash *string) (*model.Nep11TransferNotification, error) {
	return nep11TransferNotificationRepository.Nep11TransferNotificationByTransactionHash(transactionhash)
}

func (r *queryResolver) ScCallsByContractHashAddress(ctx context.Context, contracthash *string, address *string) ([]*model.ScCall, error) {
	return scCallRepository.ScCallsByContractHashAddress(contracthash, address)
}

func (r *queryResolver) ScVoteCallByTransactionHash(ctx context.Context, transactionhash *string) ([]*model.ScVoteCall, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScVoteCallsByVoterAddress(ctx context.Context, voteraddress *string) ([]*model.ScVoteCall, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScVoteCallsByCandidateAddress(ctx context.Context, candidateaddress *string) ([]*model.ScVoteCall, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScVoteCall(ctx context.Context) (*model.ScVoteCall, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VotersByCandidateAddress(ctx context.Context, candidateaddress *string) ([]*model.Vote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VotesByCandidateAddress(ctx context.Context, candidateaddress *string) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionsByAddress(ctx context.Context, address *string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionsByBlockHash(ctx context.Context, address *string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionsByBlockHeight(ctx context.Context, height *int) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionsBySender(ctx context.Context, address *string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionByTransactionHash(ctx context.Context, transactionhash *string) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Transaction(ctx context.Context) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransferNotificationsByAddress(ctx context.Context, address *string) ([]*model.TransferNotification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransferNotificationsByContractHash(ctx context.Context, contracthash *string) ([]*model.TransferNotification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransferNotificationByTransactionHash(ctx context.Context, transactionhash *string) (*model.TransferNotification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransferNotification(ctx context.Context) (*model.TransferNotification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Votes(ctx context.Context) ([]*model.Vote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Vote(ctx context.Context) (*model.Vote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Candidates(ctx context.Context) ([]*model.Candidate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Candidate(ctx context.Context) (*model.Candidate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Assets(ctx context.Context) ([]*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Asset(ctx context.Context) (*model.Asset, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Headers(ctx context.Context) ([]*model.Header, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Header(ctx context.Context) (*model.Header, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	blockRepository                     repository.BlockRepository                     = repository.NewBlockRepository()
	nep11TransferNotificationRepository repository.Nep11TransferNotificationRepository = repository.NewNep11TransferNotificationRepository()
	scCallRepository                    repository.ScCallRepository                    = repository.NewScCallRepository()
	scVoteRepository                    repository.ScVoteCallRepository                = repository.NewScVoteCallRepository()
	//assetRepository repository.AssetRepository = repository.NewAssetRepository()
)
