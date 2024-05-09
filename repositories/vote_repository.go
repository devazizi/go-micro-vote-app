package repositories

import "voteapp/entity"

type VoteRepositoryInterface interface {
	CreateVote(vote entity.Vote) (entity.Vote, error)
	IndexVote() ([]entity.Vote, error)
	// GetVoteByAnalytics(voteId int)
}

func (i *Intractor) CreateVote(vote entity.Vote) (entity.Vote, error) {
	return entity.Vote{}, nil
}

func IndexVote() ([]entity.Vote, error) {
	return []entity.Vote{
		{},
		{},
	}, nil
}
