package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// PullRequestComment represents PullRequest comment author relation with
// PullRequest author - pull request comment is a different entity than issue
// comment. Issue comment is for both PRs and Issues, PR Comment is for
// PR code snippet comments.
type PullRequestComment struct {
	RequesterID       int64
	PullRequestID     int64
	CommenterID       int64
	CommenterUserName string
	Weight            int64
}

// CreatePullRequestComment will return PullRequestComment model object
func CreatePullRequestComment(prc *github.PullRequestComment, pullRequestID, requesterID, weight int64) PullRequestComment {
	return PullRequestComment{
		RequesterID:       requesterID,
		PullRequestID:     pullRequestID,
		CommenterID:       prc.GetUser().GetID(),
		CommenterUserName: prc.GetUser().GetLogin(),
		Weight:            weight,
	}
}

// ID returns IssueComment CommenterID
func (prc PullRequestComment) ID() int64 {
	return prc.CommenterID
}

// Neo returns PullRequestComment neo.Query relation with RequesterID
func (prc PullRequestComment) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (requester:User {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:COMMENTED_PULL_REQUEST {%s: %d, PullRequestID: %d}]-(requester)
		`,
		prc.RequesterID, prc.CommenterID, prc.CommenterUserName, neo.WEIGHT_LABEL, prc.Weight, prc.PullRequestID)
	return neo.Query(queryString)
}
