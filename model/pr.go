package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type PullRequest struct {
	RepositoryID int64
	UserID       int64
	UserName     string
}

func CreatePullRequest(pr *github.PullRequest, repoID int64) PullRequest {
	return PullRequest{
		RepositoryID: repoID,
		UserID:       pr.GetUser().GetID(),
		UserName:     pr.GetUser().GetLogin(),
	}
}

func (pr PullRequest) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (repo:Repository {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[req:REQUESTED]->(repo)
		`,
		pr.RepositoryID, pr.UserID, pr.UserName)
	return neo.Query(queryString)
}
