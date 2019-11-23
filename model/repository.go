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
	OwnerID     int64
	Owner       string
	Archived    bool
}

func CreateRepository(ghRepository *github.Repository) Repository {

	return Repository{
		ID:          ghRepository.GetID(),
		Name:        ghRepository.GetName(),
		Description: ghRepository.GetDescription(),
		OwnerID:     ghRepository.GetOwner().GetID(),
		Owner:       ghRepository.GetOwner().GetLogin(),
		Archived:    ghRepository.GetArchived(),
	}
}

func (r Repository) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE (repo:Repository {
				ID: %d,
				Name: "%s",
				Description: "%s",
				Archived: %t
			})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:OWNER]-(repo)
		`, r.ID, r.Name, r.Description, r.Archived, r.OwnerID, r.Owner)
	return neo.Query(queryString)
}
