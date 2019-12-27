package ghapi

import (
	"github.com/google/go-github/github"
)

type PullRequestServicesWeights struct {
	ReviewersWeight, IssueCommentsWeight, PRCommentsWeight int64
}

func PullRequestServices(githubClient *github.Client, weights PullRequestServicesWeights) []PullRequestService {
	return []PullRequestService{
		ReviewersService{
			githubClient: githubClient,
			weight:       weights.ReviewersWeight,
		},
		IssueCommentsService{
			githubClient: githubClient,
			weight:       weights.IssueCommentsWeight,
		},
		PRCommentsService{
			githubClient: githubClient,
			weight:       weights.PRCommentsWeight,
		},
	}
}
