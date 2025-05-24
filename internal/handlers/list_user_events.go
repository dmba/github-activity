package handlers

import (
	"context"
	"fmt"
	"github-activity/internal/appcontext"
	"github-activity/pkg/github"
	"log"
)

func ListUserEvents(ctx context.Context, username string) error {
	client := appcontext.ApiClientFromContext(ctx)

	response, err := client.GetUsersEvents(ctx, username, nil)
	if err != nil {
		log.Fatal("Error fetching user events:", err)
	}

	printStatus(response)

	return nil
}

func printStatus(response *[]github.Event) {
	for _, event := range *response {
		description := getShortEventDescription(event)
		fmt.Println(description)
	}
}

func getShortEventDescription(event github.Event) string {
	switch event.Type {
	case github.CreateEvent:
		return fmt.Sprintf("- Created %s", event.Repo.Name)
	case github.DeleteEvent:
		return fmt.Sprintf("- Deleted %s", event.Repo.Name)
	case github.IssueCommentEvent:
		return fmt.Sprintf("- Commented on issue in %s", event.Repo.Name)
	case github.IssuesEvent:
		return fmt.Sprintf("- Opened a new issue in %s", event.Repo.Name)
	case github.ForkEvent:
		return fmt.Sprintf("- Forked %s", event.Repo.Name)
	case github.PublicEvent:
		return fmt.Sprintf("- Made repo %s public", event.Repo.Name)
	case github.PullRequestEvent:
		return fmt.Sprintf("- Opened Pull Request in %s", event.Repo.Name)
	case github.PullRequestReviewEvent:
		return fmt.Sprintf("- Reviewed Pull Request in %s", event.Repo.Name)
	case github.PullRequestReviewCommentEvent:
		return fmt.Sprintf("- Commented on Pull Request in %s", event.Repo.Name)
	case github.PushEvent:
		return fmt.Sprintf("- Pushed %d commits to %s", len(event.Payload.Commits), event.Repo.Name)
	case github.ReleaseEvent:
		return fmt.Sprintf("- Released %s", event.Repo.Name)
	case github.WatchEvent:
		return fmt.Sprintf("- Started watching %s", event.Repo.Name)
	default:
		return fmt.Sprintf("- Unknown event type <%s> on repo %s", event.Type, event.Repo.Name)
	}
}
