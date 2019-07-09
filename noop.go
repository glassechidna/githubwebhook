package githubwebhook

import (
	"context"
	"github.com/google/go-github/v26/github"
)

type NoopHandler struct {
}

var _ = Handler(&NoopHandler{})

func (h *NoopHandler) CheckRun(ctx context.Context, event *github.CheckRunEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) CheckSuite(ctx context.Context, event *github.CheckSuiteEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) CommitComment(ctx context.Context, event *github.CommitCommentEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Create(ctx context.Context, event *github.CreateEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Delete(ctx context.Context, event *github.DeleteEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Deployment(ctx context.Context, event *github.DeploymentEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) DeploymentStatus(ctx context.Context, event *github.DeploymentStatusEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Fork(ctx context.Context, event *github.ForkEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Gollum(ctx context.Context, event *github.GollumEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Installation(ctx context.Context, event *github.InstallationEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) InstallationRepositories(ctx context.Context, event *github.InstallationRepositoriesEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) IssueComment(ctx context.Context, event *github.IssueCommentEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Issues(ctx context.Context, event *github.IssuesEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Label(ctx context.Context, event *github.LabelEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) MarketplacePurchase(ctx context.Context, event *github.MarketplacePurchaseEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Member(ctx context.Context, event *github.MemberEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Membership(ctx context.Context, event *github.MembershipEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Milestone(ctx context.Context, event *github.MilestoneEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Organization(ctx context.Context, event *github.OrganizationEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) OrgBlock(ctx context.Context, event *github.OrgBlockEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) PageBuild(ctx context.Context, event *github.PageBuildEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Ping(ctx context.Context, event *github.PingEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Project(ctx context.Context, event *github.ProjectEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) ProjectCard(ctx context.Context, event *github.ProjectCardEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) ProjectColumn(ctx context.Context, event *github.ProjectColumnEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Public(ctx context.Context, event *github.PublicEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) PullRequest(ctx context.Context, event *github.PullRequestEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) PullRequestReview(ctx context.Context, event *github.PullRequestReviewEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) PullRequestReviewComment(ctx context.Context, event *github.PullRequestReviewCommentEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Push(ctx context.Context, event *github.PushEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Release(ctx context.Context, event *github.ReleaseEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Repository(ctx context.Context, event *github.RepositoryEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Status(ctx context.Context, event *github.StatusEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Team(ctx context.Context, event *github.TeamEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) TeamAdd(ctx context.Context, event *github.TeamAddEvent) ([]byte, error) {
	return nil, nil
}

func (h *NoopHandler) Watch(ctx context.Context, event *github.WatchEvent) ([]byte, error) {
	return nil, nil
}
