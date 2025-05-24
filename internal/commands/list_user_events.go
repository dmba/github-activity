package commands

import (
	"context"
	"github.com/urfave/cli/v3"
)

func newListUserEventsCommand(handler UserEventsHandler) *cli.Command {
	return &cli.Command{
		Name:      "github-activity",
		Usage:     "List GitHub user events",
		UsageText: "github-activity [username]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "username",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			username := cmd.StringArg("username")
			if username == "" {
				return cli.Exit("username is required", 1)
			}
			return handler(ctx, username)
		},
	}
}
