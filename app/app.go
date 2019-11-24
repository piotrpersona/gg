package app

import (
	"os"
	"sync"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/ghapi"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"

	log "github.com/sirupsen/logrus"
)

func fetchResources(
	neocfg neo.Config,
	githubClient *github.Client,
	repoResources []ghapi.RepoResource,
	repoModel model.Repository,
) {
	var resourcesWg sync.WaitGroup
	numberOfResourcesTasks := len(repoResources)
	resourcesWg.Add(numberOfResourcesTasks)
	for _, repoResource := range repoResources {
		go func(wg *sync.WaitGroup, repoResource ghapi.RepoResource) {
			defer wg.Done()
			resources, err := repoResource.Fetch(githubClient, repoModel)
			if err != nil {
				log.Error(err)
			}
			neo.Neoize(neocfg, resources...)
		}(&resourcesWg, repoResource)
	}
	resourcesWg.Wait()
}

func lastSeenRepoID(since int64, neoconfig neo.Config) (lastID int64, err error) {
	lastID = since
	if since < 0 {
		lastID, err = neo.FetchLastSeenID(neoconfig)
		return
	}
	return
}

// Run will run application with provided application config
func Run(appConfig ApplicationConfig) {
	configureLogging(appConfig.LogLevel)
	neoconfig := appConfig.neoconfig()
	githubClient := ghapi.AuthenticatedClient(appConfig.Token)

	since, err := lastSeenRepoID(appConfig.Since, neoconfig)
	if err != nil {
		log.Fatal("Unable to fetch last seen ID")
		log.Fatal(err)
		os.Exit(1)
	}
	repositories, err := ghapi.FetchRepositories(githubClient, since)
	if err != nil {
		log.Fatal("Unable to fetch Github repositories")
		log.Fatal(err)
		os.Exit(1)
	}
	neo.Neoize(neoconfig, repositories...)

	repoResources := []ghapi.RepoResource{
		ghapi.Contributors{},
		ghapi.PullRequests{},
		ghapi.Issues{},
	}

	var repoWg sync.WaitGroup
	numberOfRepositories := len(repositories)
	repoWg.Add(numberOfRepositories)
	for _, repository := range repositories {
		go func(repoWg *sync.WaitGroup, repository neo.Resource) {
			defer repoWg.Done()
			repoModel := repository.(model.Repository)
			fetchResources(
				neoconfig,
				githubClient,
				repoResources,
				repoModel,
			)
		}(&repoWg, repository)
	}
	repoWg.Wait()
}
