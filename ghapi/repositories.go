package ghapi

import (
	"context"

	"github.com/piotrpersona/gg/model"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// FetchRepositories will fetch GitHub API Repositories and map them onto
// []model.Repository.
func FetchRepositories(githubClient *github.Client, since int64) (repositories []neo.Resource, err error) {
	ctx := context.Background()
	options := &github.RepositoryListAllOptions{Since: since}
	githubRepositories, _, err := githubClient.Repositories.ListAll(ctx, options)
	if err != nil {
		return
	}
	for _, githubRepository := range githubRepositories {
		repository := model.CreateRepository(githubRepository)
		repositories = append(repositories, repository)
	}
	return repositories, nil
}
