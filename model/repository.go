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
	Owner       string
	Archived    bool
}

func CreateRepository(ghRepository *github.Repository) Repository {
	return Repository{
		ID:          ghRepository.GetID(),
		Name:        ghRepository.GetName(),
		Description: ghRepository.GetDescription(),
		Owner:       ghRepository.GetOwner().GetLogin(),
		Archived:    ghRepository.GetArchived(),
	}
}

func (r Repository) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE
			(node:Repository {
				ID: %d,
				Name: "%s",
				Description: "%s",
				Owner: "%s",
				Archived: %t
			})`, r.ID, r.Name, r.Description, r.Owner, r.Archived)
	return neo.Query(queryString)
}
