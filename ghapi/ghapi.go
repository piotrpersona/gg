package ghapi

import (
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// RepoResource represents repository related GitHub resource
type RepoResource interface {
	FetchRepoResource(model.Repository) ([]neo.Resource, error)
}

type PullRequestService interface {
	Fetch(repo model.Repository, pullRequestID, requesterID int64) ([]neo.Resource, error)
}
