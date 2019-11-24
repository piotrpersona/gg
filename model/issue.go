package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Issue represents repository issue user as a graph relation.
type Issue struct {
	RepositoryID int64
	UserID       int64
	UserName     string
}

// CreateIssue will create model Issue object from GitHub API Issue.
func CreateIssue(i *github.Issue) Issue {
	return Issue{
		RepositoryID: i.GetRepository().GetID(),
		UserID:       i.GetUser().GetID(),
		UserName:     i.GetUser().GetLogin(),
	}
}

// Neo is an implementation of neo.Resource interface.
// It will return Issue as neo4j query string.
// There will be created Repository and issuer user nodes.
// Repository will be connected with issues node with relation ISSUER.
func (i Issue) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[req:ISSUER]-(repo)
		`,
		i.RepositoryID, i.UserID, i.UserName)
	return neo.Query(queryString)
}
