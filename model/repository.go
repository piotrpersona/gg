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
	Owner       string
	OwnerURL    string
	Archived    bool
}

func CreateRepository(ghRepository *github.Repository) Repository {
	return Repository{
		ID:          ghRepository.GetID(),
		Name:        ghRepository.GetName(),
		Description: ghRepository.GetDescription(),
		URL:         ghRepository.GetURL(),
		Owner:       ghRepository.GetOwner().GetLogin(),
		OwnerURL:    ghRepository.GetOwner().GetURL(),
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
				URL: "%s",
				Owner: "%s",
				OwnerURL: "%s",
				Archived: %t
			})`, r.ID, r.Name, r.Description, r.URL, r.Owner, r.OwnerURL, r.Archived)
	return neo.Query(queryString)
}
