package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

func FetchPullRequests(githubClient *github.Client, r model.Repository) (pullRequests []neo.Resource, err error) {
	ctx := context.Background()
	repoOwner := r.Owner
	repo := r.Name
	repoID := r.ID
	options := &github.PullRequestListOptions{}
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
