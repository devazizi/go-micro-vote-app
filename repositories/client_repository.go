package repositories

import "voteapp/entity"

type ClientRepositoryInterface interface {
	CreateClient(client entity.Client) (entity.Client, error)
	FindClient(clientId int) (entity.Client, error)
}

func (i Intractor) CreateClient(client entity.Client) (entity.Client, error) {

	return entity.Client{}, nil
}

func (i Intractor) FindClient(clientId int) (entity.Client, error) {

	return entity.Client{}, nil
}
