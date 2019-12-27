package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// RequestersService is responsible for mapping github.PullRequest.User to
// model.Requester
type RequestersService struct {
	GithubClient *github.Client
}

// FetchRepoResource will download PullRequest Authors
func (rs RequestersService) FetchRepoResource(repo model.Repository) (requesters []neo.Resource, err error) {
	ctx := context.Background()
	options := github.PullRequestListOptions{}
	pullRequests, _, err := rs.GithubClient.PullRequests.List(ctx, repo.Owner, repo.Name, &options)
	if err != nil {
		return
	}
	for _, pullRequest := range pullRequests {
		requester := model.CreateRequester(pullRequest)
		requesters = append(requesters, requester)
	}
	return
}
