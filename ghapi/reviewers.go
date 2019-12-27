package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// ReviewersService is responsible for mapping github.Reviewer.User to
// model.Reviewer
type ReviewersService struct {
	githubClient *github.Client
	weight       int64
}

// Fetch will download Review Authors
func (rs ReviewersService) Fetch(repo model.Repository, pullRequestID, requesterID int64) (reviewers []neo.Resource, err error) {
	ctx := context.Background()
	options := github.ListOptions{}
	githubReviewers, _, err := rs.githubClient.PullRequests.ListReviews(
		ctx, repo.Owner, repo.Name, int(pullRequestID), &options)
	if err != nil {
		return
	}
	for _, githubReviewer := range githubReviewers {
		reviewer := model.CreateReviewer(githubReviewer, pullRequestID, requesterID, rs.weight)
		if reviewer.ID() == requesterID {
			continue
		}
		reviewers = append(reviewers, reviewer)
	}
	return
}
