package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type user struct {
	Name, URL, Company string
	ID                 int64
	Followers          []*user
	Following          []*user
}

func (u user) String() string {
	var followersString, followingString string
	if u.Followers != nil {
		followersArr := make([]string, len(u.Followers))
		for index, follower := range u.Followers {
			followersArr[index] = follower.String()
		}
		followersString = strings.Join(followersArr, ", ")
	}
	if u.Following != nil {
		followingArr := make([]string, len(u.Following))
		for index, following := range u.Following {
			followingArr[index] = following.String()
		}
		followingString = strings.Join(followingArr, ", ")
	}
	baseInfo := fmt.Sprintf("Name: %s, URL: %s, Company: %s, ID: %d", u.Name, u.URL, u.Company, u.ID)
	if followersString != "" {
		baseInfo += fmt.Sprintf("\nFollowers: %s", followersString)
	}
	if followingString != "" {
		baseInfo += fmt.Sprintf("\nFollowing: %s", followingString)
	}
	return baseInfo
}

func getFollowers(ghClient *github.Client, username string) ([]*github.User, error) {
	ctx := context.Background()
	followers, _, err := ghClient.Users.ListFollowers(ctx, username, nil)
	return followers, err
}

func getFollowing(ghClient *github.Client, username string) ([]*github.User, error) {
	ctx := context.Background()
	following, _, err := ghClient.Users.ListFollowing(ctx, username, nil)
	return following, err
}

func userLeaf(githubUser *github.User) *user {
	return &user{
		Name:      githubUser.GetLogin(),
		ID:        githubUser.GetID(),
		URL:       githubUser.GetURL(),
		Company:   githubUser.GetCompany(),
		Followers: nil,
		Following: nil,
	}
}

func createFellas(githubFollow []*github.User) []*user {
	followUsers := make([]*user, len(githubFollow))
	for index, user := range githubFollow {
		followUser := userLeaf(user)
		followUsers[index] = followUser
	}
	return followUsers
}

func CreateUser(ghClient *github.Client, username string) (userRecord *user, err error) {
	ctx := context.Background()
	userData, _, err := ghClient.Users.Get(ctx, username)
	if err != nil {
		return
	}
	followers, err := getFollowers(ghClient, username)
	if err != nil {
		return
	}
	following, err := getFollowing(ghClient, username)
	if err != nil {
		return
	}
	userRecord = &user{
		Name:      userData.GetName(),
		ID:        userData.GetID(),
		URL:       userData.GetURL(),
		Company:   userData.GetCompany(),
		Followers: createFellas(followers),
		Following: createFellas(following),
	}
	return
}
