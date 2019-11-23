package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type Contributor struct {
	ID           int64
	Name         string
	URL          string
	RepositoryID int64
	ReposURL     string
}

func CreateContirbutor(ghContirbutor *github.Contributor, repoID int64) Contributor {
	return Contributor{
		ID:           ghContirbutor.GetID(),
		Name:         ghContirbutor.GetLogin(),
		URL:          ghContirbutor.GetURL(),
		RepositoryID: repoID,
		ReposURL:     ghContirbutor.GetReposURL(),
	}
}

func (c Contributor) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (node:User:Contirbutor {
			ID: %d,
			Name: "%s",
			URL: "%s",
			ReposURL: "%s"
		})
		MERGE (node)-[r:CONTRIBUTES_TO]->(repo)`,
		c.RepositoryID, c.ID, c.Name, c.URL, c.ReposURL)
	return neo.Query(queryString)
}
