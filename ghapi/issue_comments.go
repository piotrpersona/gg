package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

type IssueCommentsService struct {
	githubClient *github.Client
	weight       int64
}

func (ics IssueCommentsService) Fetch(repo model.Repository, pullRequestID, requesterID int64) (issueComments []neo.Resource, err error) {
	ctx := context.Background()
	options := github.IssueListCommentsOptions{}
	githubIssueComments, _, err := ics.githubClient.Issues.ListComments(
		ctx, repo.Owner, repo.Name, int(pullRequestID), &options)
	if err != nil {
		return
	}
	for _, githubIssueComment := range githubIssueComments {
		issueComment := model.CreateIssueComment(githubIssueComment, pullRequestID, requesterID, ics.weight)
		if issueComment.ID() == requesterID {
			continue
		}
		issueComments = append(issueComments, issueComment)
	}
	return
}
