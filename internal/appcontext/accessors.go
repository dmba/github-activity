package appcontext

import (
	"context"
	"github-activity/pkg/github"
	"log"
)

func ApiClientFromContext(ctx context.Context) *github.APIClient {
	if l, ok := ctx.Value(ctxKeyGithubApiClient).(*github.APIClient); ok {
		return l
	}
	log.Fatal("APIClient not found in context")
	return nil
}
