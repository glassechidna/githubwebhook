package githubwebhook

import (
	"context"
	"github.com/google/go-github/v26/github"
	"io/ioutil"
	"net/http"
)

type Http struct {
	Handler   Handler
	SecretKey []byte
	Contexter func(ctx context.Context, event interface{}) context.Context
}

func (h *Http) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload []byte
	var err error

	if h.SecretKey != nil && len(h.SecretKey) > 0 {
		payload, err = github.ValidatePayload(r, h.SecretKey)
	} else {
		payload, err = ioutil.ReadAll(r.Body)
	}
	if err != nil {
		panic(err)
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		panic(err)
	}

	var resp []byte

	ctx := r.Context()
	if h.Contexter != nil {
		ctx = h.Contexter(ctx, event)
	}

	switch event := event.(type) {
	case *github.CheckRunEvent:
		resp, err = h.Handler.CheckRun(ctx, event)
	case *github.CheckSuiteEvent:
		resp, err = h.Handler.CheckSuite(ctx, event)
	case *github.CommitCommentEvent:
		resp, err = h.Handler.CommitComment(ctx, event)
	case *github.CreateEvent:
		resp, err = h.Handler.Create(ctx, event)
	case *github.DeleteEvent:
		resp, err = h.Handler.Delete(ctx, event)
	case *github.DeploymentEvent:
		resp, err = h.Handler.Deployment(ctx, event)
	case *github.DeploymentStatusEvent:
		resp, err = h.Handler.DeploymentStatus(ctx, event)
	case *github.ForkEvent:
		resp, err = h.Handler.Fork(ctx, event)
	case *github.GollumEvent:
		resp, err = h.Handler.Gollum(ctx, event)
	case *github.InstallationEvent:
		resp, err = h.Handler.Installation(ctx, event)
	case *github.InstallationRepositoriesEvent:
		resp, err = h.Handler.InstallationRepositories(ctx, event)
	case *github.IssueCommentEvent:
		resp, err = h.Handler.IssueComment(ctx, event)
	case *github.IssuesEvent:
		resp, err = h.Handler.Issues(ctx, event)
	case *github.LabelEvent:
		resp, err = h.Handler.Label(ctx, event)
	case *github.MarketplacePurchaseEvent:
		resp, err = h.Handler.MarketplacePurchase(ctx, event)
	case *github.MemberEvent:
		resp, err = h.Handler.Member(ctx, event)
	case *github.MembershipEvent:
		resp, err = h.Handler.Membership(ctx, event)
	case *github.MilestoneEvent:
		resp, err = h.Handler.Milestone(ctx, event)
	case *github.OrganizationEvent:
		resp, err = h.Handler.Organization(ctx, event)
	case *github.OrgBlockEvent:
		resp, err = h.Handler.OrgBlock(ctx, event)
	case *github.PageBuildEvent:
		resp, err = h.Handler.PageBuild(ctx, event)
	case *github.PingEvent:
		resp, err = h.Handler.Ping(ctx, event)
	case *github.ProjectEvent:
		resp, err = h.Handler.Project(ctx, event)
	case *github.ProjectCardEvent:
		resp, err = h.Handler.ProjectCard(ctx, event)
	case *github.ProjectColumnEvent:
		resp, err = h.Handler.ProjectColumn(ctx, event)
	case *github.PublicEvent:
		resp, err = h.Handler.Public(ctx, event)
	case *github.PullRequestEvent:
		resp, err = h.Handler.PullRequest(ctx, event)
	case *github.PullRequestReviewEvent:
		resp, err = h.Handler.PullRequestReview(ctx, event)
	case *github.PullRequestReviewCommentEvent:
		resp, err = h.Handler.PullRequestReviewComment(ctx, event)
	case *github.PushEvent:
		resp, err = h.Handler.Push(ctx, event)
	case *github.ReleaseEvent:
		resp, err = h.Handler.Release(ctx, event)
	case *github.RepositoryEvent:
		resp, err = h.Handler.Repository(ctx, event)
	case *github.StatusEvent:
		resp, err = h.Handler.Status(ctx, event)
	case *github.TeamEvent:
		resp, err = h.Handler.Team(ctx, event)
	case *github.TeamAddEvent:
		resp, err = h.Handler.TeamAdd(ctx, event)
	case *github.WatchEvent:
		resp, err = h.Handler.Watch(ctx, event)
	}

	if err != nil {
		panic(err)
	}

	w.Write(resp)
}
