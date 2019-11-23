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

func CreateContirbutor(ghContirbutor *github.Contributor, repoID int64) Contributor {
	return Contributor{
		ID:           ghContirbutor.GetID(),
		Name:         ghContirbutor.GetLogin(),
		RepositoryID: repoID,
	}
}

func (c Contributor) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (node:User:Contirbutor {
			ID: %d,
			Name: "%s"
		})
		MERGE (node)-[r:CONTRIBUTES_TO]->(repo)`,
		c.RepositoryID, c.ID, c.Name)
	return neo.Query(queryString)
}
