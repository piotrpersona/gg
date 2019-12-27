package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type Requester struct {
	PullRequestID int64
	RequesterID   int64
	UserName      string
}

func CreateRequester(pr *github.PullRequest) Requester {
	return Requester{
		PullRequestID: int64(pr.GetNumber()),
		RequesterID:   pr.GetUser().GetID(),
		UserName:      pr.GetUser().GetLogin(),
	}
}

func (r Requester) ID() int64 {
	return r.RequesterID
}

func (r Requester) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE (user:User {
			ID: %d,
			Name: "%s"
		})`,
		r.RequesterID, r.UserName)
	return neo.Query(queryString)
}
