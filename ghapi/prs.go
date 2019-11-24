package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// Contributors represents GitHub API PullRequests mapping with model resource.
type PullRequests struct{}

// Fetch is an implementation of RepoResource interface.
// It will create []model.PullRequest from GitHub API PullRequest.
func (pr PullRequests) Fetch(githubClient *github.Client, r model.Repository) (pullRequests []neo.Resource, err error) {
	ctx := context.Background()
	repoOwner := r.Owner
	repo := r.Name
	repoID := r.ID
	options := &github.PullRequestListOptions{State: "all"}
	githubPullRequests, _, err := githubClient.PullRequests.List(ctx, repoOwner, repo, options)
	if err != nil {
		return
	}
	for _, githubPullRequest := range githubPullRequests {
		pullRequest := model.CreatePullRequest(githubPullRequest, repoID)
		pullRequests = append(pullRequests, pullRequest)
	}
	return
}
