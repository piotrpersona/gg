package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// Contributors represents GitHub API Contributors mapping with model resource.
type Contributors struct{}

// Fetch is an implementation of RepoResource interface.
// It will create []model.Contributor from GitHub API Contributors.
func (c Contributors) Fetch(githubClient *github.Client, r model.Repository) (contributors []neo.Resource, err error) {
	ctx := context.Background()
	repoOwner := r.Owner
	repo := r.Name
	repoID := r.ID
	options := &github.ListContributorsOptions{}
	githubContributors, _, err := githubClient.Repositories.ListContributors(ctx, repoOwner, repo, options)
	if err != nil {
		return
	}
	for _, githubContributor := range githubContributors {
		contirbutor := model.CreateContributor(githubContributor, repoID)
		contributors = append(contributors, contirbutor)
	}
	return
}
