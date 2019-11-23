package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

func FetchIssues(githubClient *github.Client, r model.Repository) (issues []neo.Resource, err error) {
	ctx := context.Background()
	repoOwner := r.Owner
	repo := r.Name
	options := &github.IssueListByRepoOptions{}
	githubIssues, _, err := githubClient.Issues.ListByRepo(ctx, repoOwner, repo, options)
	if err != nil {
		return
	}
	for _, githubIssue := range githubIssues {
		issue := model.CreateIssue(githubIssue)
		issues = append(issues, issue)
	}
	return
}
