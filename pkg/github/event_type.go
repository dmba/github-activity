package github

type EventType string

const (
	CreateEvent                   EventType = "CreateEvent"
	DeleteEvent                   EventType = "DeleteEvent"
	IssuesEvent                   EventType = "IssuesEvent"
	IssueCommentEvent             EventType = "IssueCommentEvent"
	ForkEvent                     EventType = "ForkEvent"
	PublicEvent                   EventType = "PublicEvent"
	PullRequestEvent              EventType = "PullRequestEvent"
	PullRequestReviewEvent        EventType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventType = "PullRequestReviewCommentEvent"
	PushEvent                     EventType = "PushEvent"
	ReleaseEvent                  EventType = "ReleaseEvent"
	WatchEvent                    EventType = "WatchEvent"
)
