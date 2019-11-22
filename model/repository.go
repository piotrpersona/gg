package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Repository represents github repository node
type Repository struct {
	ID          int64
	Name        string
	Description string
	Archived    bool
}

func CreateRepository(ghRepository *github.Repository) Repository {
	return Repository{
		ID:          ghRepository.GetID(),
		Name:        ghRepository.GetName(),
		Description: ghRepository.GetDescription(),
		Archived:    ghRepository.GetArchived(),
	}
}

func (r Repository) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE
			(node:repository {
				ID: %d,
				Name: "%s",
				Description: "%s",
				Archived: %t
			})`, r.ID, r.Name, r.Description, r.Archived)
	return neo.Query(queryString)
}
