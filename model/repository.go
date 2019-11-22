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
	URL         string
	Archived    bool
}

func CreateRepository(ghRepository *github.Repository) Repository {
	return Repository{
		ID:          ghRepository.GetID(),
		Name:        ghRepository.GetName(),
		Description: ghRepository.GetDescription(),
		URL:         ghRepository.GetURL(),
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
				URL: "%s",
				Archived: %t
			})`, r.ID, r.Name, r.Description, r.URL, r.Archived)
	return neo.Query(queryString)
}
