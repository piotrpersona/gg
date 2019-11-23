package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type Issue struct {
	RepositoryID int64
	UserID       int64
	UserName     string
}

func CreateIssue(i *github.Issue) Issue {
	return Issue{
		RepositoryID: i.GetRepository().GetID(),
		UserID:       i.GetUser().GetID(),
		UserName:     i.GetUser().GetLogin(),
	}
}

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
