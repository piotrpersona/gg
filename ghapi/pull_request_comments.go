package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

type PRCommentsService struct {
	githubClient *github.Client
	weight       int64
}

func (prcs PRCommentsService) Fetch(repo model.Repository, pullRequestID, requesterID int64) (prComments []neo.Resource, err error) {
	ctx := context.Background()
	options := github.PullRequestListCommentsOptions{}
	githubPRComments, _, err := prcs.githubClient.PullRequests.ListComments(
		ctx, repo.Owner, repo.Name, int(pullRequestID), &options)
	if err != nil {
		return
	}
	for _, githubPrComment := range githubPRComments {
		prComment := model.CreatePullRequestComment(githubPrComment, pullRequestID, requesterID, prcs.weight)
		if prComment.ID() == requesterID {
			continue
		}
		prComments = append(prComments, prComment)
	}
	return
}
