package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Contributor represents repository contributor user as a graph relation.
type Contributor struct {
	ID           int64
	Name         string
	RepositoryID int64
}

// CreateContributor will create model Contributor object from GitHub API Contributor.
func CreateContributor(ghContributor *github.Contributor, repoID int64) Contributor {
	return Contributor{
		ID:           ghContributor.GetID(),
		Name:         ghContributor.GetLogin(),
		RepositoryID: repoID,
	}
}

// Neo is an implementation of neo.Resource interface.
// It will return Contributor as neo4j query string.
// There will be created Repository and contributor user node.
// Repository will be connected with contributor node with relation CONTRIBUTOR.
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
