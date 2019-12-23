package app

import "github.com/piotrpersona/gg/ghapi"

func repoResources() []ghapi.RepoResource {
	return []ghapi.RepoResource{
		ghapi.Contributors{},
		ghapi.PullRequests{},
		ghapi.Issues{},
	}
}
