package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

func FetchContributors(githubClient *github.Client, r model.Repository) (contributors []neo.Resource, err error) {
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
		contirbutor := model.CreateContirbutor(githubContributor, repoID)
		contributors = append(contributors, contirbutor)
	}
	return
}
