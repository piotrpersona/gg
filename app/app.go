package users

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func authenticatedGithubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	ghClient := github.NewClient(tc)
	return ghClient
}

func githubClient() *github.Client {
	return github.NewClient(nil)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getFollowers(ghClient *github.Client, user string) []*github.User {
	ctx := context.Background()
	followers, _, err := ghClient.Users.ListFollowers(ctx, user, nil)
	panicErr(err)
	return followers
}

func getFollowing(ghClient *github.Client, user string) []*github.User {
	ctx := context.Background()
	following, _, err := ghClient.Users.ListFollowing(ctx, user, nil)
	panicErr(err)
	return following
}

func App() {
	ghClient := githubClient()
	users := []string{"piotrpersona", "mateuszstompor", "reconndev", "filwie"}

	for _, user := range users {
		followers := getFollowers(ghClient, user)
		following := getFollowing(ghClient, user)
		fmt.Println(followers)
		fmt.Println(following)
	}
}
