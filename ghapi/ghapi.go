package ghapi

import (
	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"
)

// RepoResource represents repository related GitHub resource
type RepoResource interface {
	Fetch(*github.Client, model.Repository) ([]neo.Resource, error)
}
