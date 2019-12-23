package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// PullRequest represents repository pull request that was created by user as a graph relation.
type PullRequest struct {
	ID           int64
	RepositoryID int64
	UserID       int64
	UserName     string
}

// CreatePullRequest will create model PullRequest object from GitHub API PullRequest.
func CreatePullRequest(pr *github.PullRequest, repoID int64) PullRequest {
	return PullRequest{
		ID:           pr.GetID(),
		RepositoryID: repoID,
		UserID:       pr.GetUser().GetID(),
		UserName:     pr.GetUser().GetLogin(),
	}
}

// Neo is an implementation of neo.Resource interface.
// It will return PullRequest as neo4j query string.
// There will be created Repository and pull requester user nodes.
// Repository will be connected with pull requester node with relation PULL_REQUEST.
func (pr PullRequest) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[req:PULL_REQUEST {ID: %d}]-(repo)
		`,
		pr.RepositoryID, pr.UserID, pr.UserName, pr.ID)
	return neo.Query(queryString)
}
