package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type Contributor struct {
	ID           int64
	Name         string
	RepositoryID int64
}

func CreateContributor(ghContributor *github.Contributor, repoID int64) Contributor {
	return Contributor{
		ID:           ghContributor.GetID(),
		Name:         ghContributor.GetLogin(),
		RepositoryID: repoID,
	}
}

func (c Contributor) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (node:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (node)-[r:CONTRIBUTOR]-(repo)`,
		c.RepositoryID, c.ID, c.Name)
	return neo.Query(queryString)
}
