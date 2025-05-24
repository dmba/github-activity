package commands

import (
	"context"
	"github.com/urfave/cli/v3"
	"os"
)

type UserEventsHandler = func(ctx context.Context, username string) error

type GithubActivityCommandOptions struct {
	ListUserEvents UserEventsHandler
}

type GithubActivityCommand struct {
	*cli.Command
}

func NewGithubActivityCommand(opts *GithubActivityCommandOptions) *GithubActivityCommand {
	return &GithubActivityCommand{
		Command: newListUserEventsCommand(opts.ListUserEvents),
	}
}

func (tc *GithubActivityCommand) Exec(ctx context.Context) error {
	return tc.Run(ctx, os.Args)
}
