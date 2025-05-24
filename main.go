package main

import (
	"context"
	"github-activity/internal/appcontext"
	"github-activity/internal/commands"
	"github-activity/internal/handlers"
	"github-activity/pkg/config"
	"github-activity/pkg/github"
	"log"
)

func main() {
	app, err := createApp()
	if err != nil {
		log.Fatal(err)
	}

	cmd := commands.NewGithubActivityCommand(&commands.GithubActivityCommandOptions{
		ListUserEvents: handlers.ListUserEvents,
	})

	if err = cmd.Exec(app); err != nil {
		log.Fatal(err)
	}
}

func createApp() (*appcontext.AppContext, error) {
	ctx := context.Background()

	env, err := config.NewEnv()
	if err != nil {
		log.Fatalln("Error loading environment variables:", err)
		return nil, err
	}

	api := github.NewGitHubAPIClient(env.GithubToken)

	return appcontext.NewAppContext(ctx, api), nil
}
