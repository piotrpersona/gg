package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

type IssueComment struct {
	RequesterID    int64
	PullRequestID  int64
	IssuerID       int64
	IssuerUserName string
	Weight         int64
}

func CreateIssueComment(prr *github.IssueComment, pullRequestID, requesterID, weight int64) IssueComment {
	return IssueComment{
		RequesterID:    requesterID,
		PullRequestID:  pullRequestID,
		IssuerID:       prr.GetUser().GetID(),
		IssuerUserName: prr.GetUser().GetLogin(),
		Weight:         weight,
	}
}

func (r IssueComment) ID() int64 {
	return r.IssuerID
}

func (r IssueComment) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (requester:User {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:COMMENTED_ISSUE {Weight: %d, PullRequestID: %d}]-(requester)
		`,
		r.RequesterID, r.IssuerID, r.IssuerUserName, r.Weight, r.PullRequestID)
	return neo.Query(queryString)
}
