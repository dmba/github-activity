package appcontext

import (
	"context"
	"github-activity/pkg/github"
)

type AppContext struct {
	context.Context
}

func NewAppContext(ctx context.Context, api *github.APIClient) *AppContext {
	ctx = context.WithValue(ctx, ctxKeyGithubApiClient, api)
	return &AppContext{
		Context: ctx,
	}
}
