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

func (ic IssueComment) ID() int64 {
	return ic.IssuerID
}

func (ic IssueComment) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (requester:User {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:COMMENTED_ISSUE {%s: %d, PullRequestID: %d}]-(requester)
		`,
		ic.RequesterID, ic.IssuerID, ic.IssuerUserName, neo.WEIGHT_LABEL, ic.Weight, ic.PullRequestID)
	return neo.Query(queryString)
}
