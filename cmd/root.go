package cmd

import (
	"os"

	"github.com/piotrpersona/gg/app"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func buildRootCmd() (rootCmd *cobra.Command) {
	var (
		since                                    int64
		uri, username, password, token, loglevel string
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
				Since:    since,
				LogLevel: level,
			}
			app.Run(applicationConfig)
		},
	}

	flags := rootCmd.Flags()
	flags.Int64VarP(&since, "since", "s", -1, "Starting point of repositories to fetch")
	flags.StringVarP(&uri, "uri", "", viper.GetString("NEO_URI"), "Neo4j compatible URI")
	flags.StringVarP(&username, "username", "u", viper.GetString("NEO_USER"), "Neo4j connection username")
	flags.StringVarP(&password, "password", "p", viper.GetString("NEO_PASS"), "Neo4j connection password")
	flags.StringVarP(&token, "token", "t", viper.GetString("GITHUB_TOKEN"), "GitHub API Token String")
	flags.StringVarP(&loglevel, "loglevel", "", log.InfoLevel.String(), "Log level")

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
