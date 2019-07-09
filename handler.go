package githubwebhook

import (
	"context"
	"github.com/google/go-github/v26/github"
)

type Handler interface {
	CheckRun(ctx context.Context, event *github.CheckRunEvent) ([]byte, error)
	CheckSuite(ctx context.Context, event *github.CheckSuiteEvent) ([]byte, error)
	CommitComment(ctx context.Context, event *github.CommitCommentEvent) ([]byte, error)
	Create(ctx context.Context, event *github.CreateEvent) ([]byte, error)
	Delete(ctx context.Context, event *github.DeleteEvent) ([]byte, error)
	Deployment(ctx context.Context, event *github.DeploymentEvent) ([]byte, error)
	DeploymentStatus(ctx context.Context, event *github.DeploymentStatusEvent) ([]byte, error)
	Fork(ctx context.Context, event *github.ForkEvent) ([]byte, error)
	Gollum(ctx context.Context, event *github.GollumEvent) ([]byte, error)
	Installation(ctx context.Context, event *github.InstallationEvent) ([]byte, error)
	InstallationRepositories(ctx context.Context, event *github.InstallationRepositoriesEvent) ([]byte, error)
	IssueComment(ctx context.Context, event *github.IssueCommentEvent) ([]byte, error)
	Issues(ctx context.Context, event *github.IssuesEvent) ([]byte, error)
	Label(ctx context.Context, event *github.LabelEvent) ([]byte, error)
	MarketplacePurchase(ctx context.Context, event *github.MarketplacePurchaseEvent) ([]byte, error)
	Member(ctx context.Context, event *github.MemberEvent) ([]byte, error)
	Membership(ctx context.Context, event *github.MembershipEvent) ([]byte, error)
	Milestone(ctx context.Context, event *github.MilestoneEvent) ([]byte, error)
	Organization(ctx context.Context, event *github.OrganizationEvent) ([]byte, error)
	OrgBlock(ctx context.Context, event *github.OrgBlockEvent) ([]byte, error)
	PageBuild(ctx context.Context, event *github.PageBuildEvent) ([]byte, error)
	Ping(ctx context.Context, event *github.PingEvent) ([]byte, error)
	Project(ctx context.Context, event *github.ProjectEvent) ([]byte, error)
	ProjectCard(ctx context.Context, event *github.ProjectCardEvent) ([]byte, error)
	ProjectColumn(ctx context.Context, event *github.ProjectColumnEvent) ([]byte, error)
	Public(ctx context.Context, event *github.PublicEvent) ([]byte, error)
	PullRequest(ctx context.Context, event *github.PullRequestEvent) ([]byte, error)
	PullRequestReview(ctx context.Context, event *github.PullRequestReviewEvent) ([]byte, error)
	PullRequestReviewComment(ctx context.Context, event *github.PullRequestReviewCommentEvent) ([]byte, error)
	Push(ctx context.Context, event *github.PushEvent) ([]byte, error)
	Release(ctx context.Context, event *github.ReleaseEvent) ([]byte, error)
	Repository(ctx context.Context, event *github.RepositoryEvent) ([]byte, error)
	Status(ctx context.Context, event *github.StatusEvent) ([]byte, error)
	Team(ctx context.Context, event *github.TeamEvent) ([]byte, error)
	TeamAdd(ctx context.Context, event *github.TeamAddEvent) ([]byte, error)
	Watch(ctx context.Context, event *github.WatchEvent) ([]byte, error)
}
