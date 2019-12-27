package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Requester represents Pull Request Author node entity.
type Requester struct {
	PullRequestID int64
	RequesterID   int64
	UserName      string
}

// CreateRequester will return Requester model object
func CreateRequester(pr *github.PullRequest) Requester {
	return Requester{
		PullRequestID: int64(pr.GetNumber()),
		RequesterID:   pr.GetUser().GetID(),
		UserName:      pr.GetUser().GetLogin(),
	}
}

// ID returns Requester IssuerID
func (r Requester) ID() int64 {
	return r.RequesterID
}

// Neo returns Requester neo.Query node representation
func (r Requester) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MERGE (user:User {
			ID: %d,
			Name: "%s"
		})`,
		r.RequesterID, r.UserName)
	return neo.Query(queryString)
}
