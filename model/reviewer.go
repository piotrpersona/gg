package model

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// Reviewer represents Review author relation with
// PullRequest/Issue author.
type Reviewer struct {
	RequesterID      int64
	PullRequestID    int64
	ReviewerID       int64
	ReviewerUserName string
	Weight           int64
}

// CreateReviewer will return Reviewer model object
func CreateReviewer(prr *github.PullRequestReview, pullRequestID, requesterID, weight int64) Reviewer {
	return Reviewer{
		RequesterID:      requesterID,
		PullRequestID:    pullRequestID,
		ReviewerID:       prr.GetUser().GetID(),
		ReviewerUserName: prr.GetUser().GetLogin(),
		Weight:           weight,
	}
}

// ID returns Reviewer ReviewerID
func (r Reviewer) ID() int64 {
	return r.ReviewerID
}

// Neo returns Reviewer neo.Query relation with RequesterID
func (r Reviewer) Neo() neo.Query {
	queryString := fmt.Sprintf(
		`MATCH (requester:User {ID: %d})
		MERGE (user:User {
			ID: %d,
			Name: "%s"
		})
		MERGE (user)-[r:REVIEWED {%s: %d, PullRequestID: %d}]-(requester)
		`,
		r.RequesterID, r.ReviewerID, r.ReviewerUserName, neo.WEIGHT_LABEL, r.Weight, r.PullRequestID)
	return neo.Query(queryString)
}
