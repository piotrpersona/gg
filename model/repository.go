package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Repository represents repository as a graph node.
type Repository struct {
	RepoID   int64
	Name     string
	OwnerID  int64
	Owner    string
	Archived bool
}

// CreateRepository will create model Repository object from GitHub API Repository.
func CreateRepository(ghRepository *github.Repository) Repository {
	return Repository{
		RepoID:   ghRepository.GetID(),
		Name:     ghRepository.GetName(),
		OwnerID:  ghRepository.GetOwner().GetID(),
		Owner:    ghRepository.GetOwner().GetLogin(),
		Archived: ghRepository.GetArchived(),
	}
}

func (r Repository) ID() int64 {
	return r.RepoID
}

// Neo is an implementation of neo.Resource interface.
// It will return Repository as neo4j query string.
// There will be created Repository and owner user nodes.
// Repository will be connected with owner node with relation OWNER.
func (r Repository) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE (repo:Repository {
				ID: %d,
				Name: "%s",
				Archived: %t
			})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:OWNER]-(repo)
		`, r.RepoID, r.Name, r.Archived, r.OwnerID, r.Owner)
	return neo.Query(queryString)
}
