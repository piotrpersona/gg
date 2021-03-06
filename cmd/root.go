package cmd

import (
	"os"

	"github.com/piotrpersona/gg/model"

	"github.com/piotrpersona/gg/app"
	"github.com/piotrpersona/gg/ghapi"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func buildRootCmd() (rootCmd *cobra.Command) {
	var (
		uri, username, password, token, loglevel, githubReposQuery    string
		reviewersWeight, issueCommentWeight, pullRequestCommentWeight int64
	)

	rootCmd = &cobra.Command{
		Use:   "gg",
		Short: "Build Github Graph",
		Long:  `Fetch repositories from a github and build a graph`,
		Run: func(cmd *cobra.Command, args []string) {
			level, err := log.ParseLevel(loglevel)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			applicationConfig := app.ApplicationConfig{
				URI:      uri,
				Username: username,
				Password: password,
				Token:    token,
				LogLevel: level,
				PullRequestWeights: ghapi.PullRequestServicesWeights{
					ReviewersWeight:     reviewersWeight,
					IssueCommentsWeight: issueCommentWeight,
					PRCommentsWeight:    pullRequestCommentWeight,
				},
				Query: githubReposQuery,
			}
			app.Run(applicationConfig)
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&uri, "uri", "", viper.GetString("NEO_URI"), "Neo4j compatible URI")
	flags.StringVarP(&username, "username", "u", viper.GetString("NEO_USER"), "Neo4j connection username")
	flags.StringVarP(&password, "password", "p", viper.GetString("NEO_PASS"), "Neo4j connection password")
	flags.StringVarP(&token, "token", "t", viper.GetString("GITHUB_TOKEN"), "GitHub API Token String")
	flags.StringVarP(&loglevel, "loglevel", "", log.InfoLevel.String(), "Log level")
	flags.Int64Var(&reviewersWeight, "review", model.REVIEW_WEIGHT, "Weight of review")
	flags.Int64Var(&issueCommentWeight, "issue-comment", model.ISSUE_COMMENT_WEIGHT, "Weight of issue comment")
	flags.Int64Var(&pullRequestCommentWeight, "pr-comment", model.PULL_REQUEST_COMMENT_WEIGHT, "Weight of pull request comment")
	flags.StringVarP(&githubReposQuery, "query", "q", ghapi.DEFAULT_QUERY, "Github repositories search query")

	return
}

// Execute will execute root command.
func Execute() {
	viper.AutomaticEnv()
	rootCmd := buildRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
