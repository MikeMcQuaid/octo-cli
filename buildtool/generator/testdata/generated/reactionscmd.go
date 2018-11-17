// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type ReactionsCmd struct {
	ListForCommitComment              ReactionsListForCommitCommentCmd              `cmd:"" help:"List reactions for a commit comment - https://developer.github.com/v3/reactions/#list-reactions-for-a-commit-comment"`
	CreateForCommitComment            ReactionsCreateForCommitCommentCmd            `cmd:"" help:"Create reaction for a commit comment - https://developer.github.com/v3/reactions/#create-reaction-for-a-commit-comment"`
	ListForIssue                      ReactionsListForIssueCmd                      `cmd:"" help:"List reactions for an issue - https://developer.github.com/v3/reactions/#list-reactions-for-an-issue"`
	CreateForIssue                    ReactionsCreateForIssueCmd                    `cmd:"" help:"Create reaction for an issue - https://developer.github.com/v3/reactions/#create-reaction-for-an-issue"`
	ListForIssueComment               ReactionsListForIssueCommentCmd               `cmd:"" help:"List reactions for an issue comment - https://developer.github.com/v3/reactions/#list-reactions-for-an-issue-comment"`
	CreateForIssueComment             ReactionsCreateForIssueCommentCmd             `cmd:"" help:"Create reaction for an issue comment - https://developer.github.com/v3/reactions/#create-reaction-for-an-issue-comment"`
	ListForPullRequestReviewComment   ReactionsListForPullRequestReviewCommentCmd   `cmd:"" help:"List reactions for a pull request review comment - https://developer.github.com/v3/reactions/#list-reactions-for-a-pull-request-review-comment"`
	CreateForPullRequestReviewComment ReactionsCreateForPullRequestReviewCommentCmd `cmd:"" help:"Create reaction for a pull request review comment - https://developer.github.com/v3/reactions/#create-reaction-for-a-pull-request-review-comment"`
	ListForTeamDiscussion             ReactionsListForTeamDiscussionCmd             `cmd:"" help:"List reactions for a team discussion - https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion"`
	CreateForTeamDiscussion           ReactionsCreateForTeamDiscussionCmd           `cmd:"" help:"Create reaction for a team discussion - https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion"`
	ListForTeamDiscussionComment      ReactionsListForTeamDiscussionCommentCmd      `cmd:"" help:"List reactions for a team discussion comment - https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment"`
	CreateForTeamDiscussionComment    ReactionsCreateForTeamDiscussionCommentCmd    `cmd:"" help:"Create reaction for a team discussion comment - https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment"`
	Delete                            ReactionsDeleteCmd                            `cmd:"" help:"Delete a reaction - https://developer.github.com/v3/reactions/#delete-a-reaction"`
}

type ReactionsListForCommitCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a commit comment."`
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForCommitCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForCommitCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the commit comment."`
}

func (c *ReactionsCreateForCommitCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsListForIssueCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	Number       int64  `required:"" name:"number"`
	Content      string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to an issue."`
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForIssueCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/issues/:number/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForIssueCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	Number       int64  `required:"" name:"number"`
	Content      string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the issue."`
}

func (c *ReactionsCreateForIssueCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/issues/:number/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsListForIssueCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to an issue comment."`
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForIssueCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/issues/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForIssueCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the issue comment."`
}

func (c *ReactionsCreateForIssueCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/issues/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsListForPullRequestReviewCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a pull request review comment."`
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForPullRequestReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForPullRequestReviewCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
	Content      string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the pull request review comment."`
}

func (c *ReactionsCreateForPullRequestReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments/:comment_id/reactions")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsListForTeamDiscussionCmd struct {
	internal.BaseCmd
	Echo             bool   "name:\"echo-preview\" required:\"\" help:\"**Note:** The team discussions API is currently available for developers to preview. See the [blog post](/changes/2018-02-07-team-discussions-api) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.echo-preview+json\n\n```\""
	SquirrelGirl     bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	Content          string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion."`
	PerPage          int64  `name:"per_page" help:"Results per page (max 100)"`
	Page             int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForTeamDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/teams/:team_id/discussions/:discussion_number/reactions")
	c.UpdatePreview("echo", c.Echo)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("team_id", c.TeamId)
	c.UpdateURLPath("discussion_number", c.DiscussionNumber)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForTeamDiscussionCmd struct {
	internal.BaseCmd
	Echo             bool   "name:\"echo-preview\" required:\"\" help:\"**Note:** The team discussions API is currently available for developers to preview. See the [blog post](/changes/2018-02-07-team-discussions-api) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.echo-preview+json\n\n```\""
	SquirrelGirl     bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	Content          string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion."`
}

func (c *ReactionsCreateForTeamDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/teams/:team_id/discussions/:discussion_number/reactions")
	c.UpdatePreview("echo", c.Echo)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("team_id", c.TeamId)
	c.UpdateURLPath("discussion_number", c.DiscussionNumber)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsListForTeamDiscussionCommentCmd struct {
	internal.BaseCmd
	Echo             bool   "name:\"echo-preview\" required:\"\" help:\"**Note:** The team discussions API is currently available for developers to preview. See the [blog post](/changes/2018-02-07-team-discussions-api) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.echo-preview+json\n\n```\""
	SquirrelGirl     bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	CommentNumber    int64  `required:"" name:"comment_number"`
	Content          string `name:"content" help:"Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion comment."`
	PerPage          int64  `name:"per_page" help:"Results per page (max 100)"`
	Page             int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ReactionsListForTeamDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/teams/:team_id/discussions/:discussion_number/comments/:comment_number/reactions")
	c.UpdatePreview("echo", c.Echo)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("team_id", c.TeamId)
	c.UpdateURLPath("discussion_number", c.DiscussionNumber)
	c.UpdateURLPath("comment_number", c.CommentNumber)
	c.UpdateURLQuery("content", c.Content)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ReactionsCreateForTeamDiscussionCommentCmd struct {
	internal.BaseCmd
	Echo             bool   "name:\"echo-preview\" required:\"\" help:\"**Note:** The team discussions API is currently available for developers to preview. See the [blog post](/changes/2018-02-07-team-discussions-api) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.echo-preview+json\n\n```\""
	SquirrelGirl     bool   "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	CommentNumber    int64  `required:"" name:"comment_number"`
	Content          string `required:"" name:"content" help:"The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion comment."`
}

func (c *ReactionsCreateForTeamDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/teams/:team_id/discussions/:discussion_number/comments/:comment_number/reactions")
	c.UpdatePreview("echo", c.Echo)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("team_id", c.TeamId)
	c.UpdateURLPath("discussion_number", c.DiscussionNumber)
	c.UpdateURLPath("comment_number", c.CommentNumber)
	c.UpdateBody("content", c.Content)
	return c.DoRequest("POST")
}

type ReactionsDeleteCmd struct {
	internal.BaseCmd
	Echo         bool  "name:\"echo-preview\" required:\"\" help:\"**Note:** The team discussions API is currently available for developers to preview. See the [blog post](/changes/2018-02-07-team-discussions-api) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.echo-preview+json\n\n```\""
	SquirrelGirl bool  "name:\"squirrel-girl-preview\" required:\"\" help:\"**Note:** APIs for managing reactions are currently available for developers to preview. See the [blog post](/changes/2016-05-12-reactions-api-preview) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview+json\n\n```\""
	ReactionId   int64 `required:"" name:"reaction_id"`
}

func (c *ReactionsDeleteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/reactions/:reaction_id")
	c.UpdatePreview("echo", c.Echo)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("reaction_id", c.ReactionId)
	return c.DoRequest("DELETE")
}
