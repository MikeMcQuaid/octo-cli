// Code generated by go-github-cli/generator; DO NOT EDIT.

package services

type PullsCmd struct {
	List                 PullsListCmd                 `cmd:"" help:"List pull requests"`
	Get                  PullsGetCmd                  `cmd:"" help:"Get a single pull request"`
	Create               PullsCreateCmd               `cmd:"" help:"Create a pull request"`
	CreateFromIssue      PullsCreateFromIssueCmd      `cmd:"" help:"Create a Pull Request from an Issue"`
	Update               PullsUpdateCmd               `cmd:"" help:"Update a pull request"`
	ListCommits          PullsListCommitsCmd          `cmd:"" help:"List commits on a pull request"`
	ListFiles            PullsListFilesCmd            `cmd:"" help:"List pull requests files"`
	CheckIfMerged        PullsCheckIfMergedCmd        `cmd:"" help:"Get if a pull request has been merged"`
	Merge                PullsMergeCmd                `cmd:"" help:"Merge a pull request (Merge Button)"`
	ListReviews          PullsListReviewsCmd          `cmd:"" help:"List reviews on a pull request"`
	GetReview            PullsGetReviewCmd            `cmd:"" help:"Get a single review"`
	DeletePendingReview  PullsDeletePendingReviewCmd  `cmd:"" help:"Delete a pending review"`
	GetCommentsForReview PullsGetCommentsForReviewCmd `cmd:"" help:"Get comments for a single review"`
	SubmitReview         PullsSubmitReviewCmd         `cmd:"" help:"Submit a pull request review"`
	DismissReview        PullsDismissReviewCmd        `cmd:"" help:"Dismiss a pull request review"`
	ListComments         PullsListCommentsCmd         `cmd:"" help:"List comments on a pull request"`
	ListCommentsForRepo  PullsListCommentsForRepoCmd  `cmd:"" help:"List comments in a repository"`
	GetComment           PullsGetCommentCmd           `cmd:"" help:"Get a single comment"`
	CreateComment        PullsCreateCommentCmd        `cmd:"" help:"Create a comment"`
	CreateCommentReply   PullsCreateCommentReplyCmd   `cmd:"" help:"Create a comment reply"`
	EditComment          PullsEditCommentCmd          `cmd:"" help:"Edit a comment"`
	DeleteComment        PullsDeleteCommentCmd        `cmd:"" help:"Delete a comment"`
	ListReviewRequests   PullsListReviewRequestsCmd   `cmd:"" help:"List review requests"`
	CreateReviewRequest  PullsCreateReviewRequestCmd  `cmd:"" help:"Create a review request"`
	DeleteReviewRequest  PullsDeleteReviewRequestCmd  `cmd:"" help:"Delete a review request"`
}

type PullsListCmd struct {
	baseCmd
	Symmetra  bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	State     string "name:\"state\" help:\"Either `open`, `closed`, or `all` to filter by state.\""
	Head      string "name:\"head\" help:\"Filter pulls by head user and branch name in the format of `user:ref-name`. Example: `github:new-script-format`.\""
	Base      string "name:\"base\" help:\"Filter pulls by base branch name. Example: `gh-pages`.\""
	Sort      string "name:\"sort\" help:\"What to sort results by. Can be either `created`, `updated`, `popularity` (comment count) or `long-running` (age, filtering by pulls updated in the last month).\""
	Direction string "name:\"direction\" help:\"The direction of the sort. Can be either `asc` or `desc`.\""
	PerPage   int64  `name:"per_page" help:"Results per page (max 100)"`
	Page      int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLQuery("state", c.State)
	c.updateURLQuery("head", c.Head)
	c.updateURLQuery("base", c.Base)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("direction", c.Direction)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsGetCmd struct {
	baseCmd
	Symmetra bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
}

func (c *PullsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	return c.doRequest("GET")
}

type PullsCreateCmd struct {
	baseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `required:"" name:"owner"`
	Repo                string `required:"" name:"repo"`
	Title               string `required:"" name:"title" help:"The title of the pull request."`
	Head                string "required:\"\" name:\"head\" help:\"The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.\""
	Base                string `required:"" name:"base" help:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository."`
	Body                string `name:"body" help:"The contents of the pull request."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateBody("title", c.Title)
	c.updateBody("head", c.Head)
	c.updateBody("base", c.Base)
	c.updateBody("body", c.Body)
	c.updateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.doRequest("POST")
}

type PullsCreateFromIssueCmd struct {
	baseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `required:"" name:"owner"`
	Repo                string `required:"" name:"repo"`
	Issue               int64  `required:"" name:"issue" help:"The issue number in this repository to turn into a Pull Request."`
	Head                string "required:\"\" name:\"head\" help:\"The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.\""
	Base                string `required:"" name:"base" help:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsCreateFromIssueCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateBody("issue", c.Issue)
	c.updateBody("head", c.Head)
	c.updateBody("base", c.Base)
	c.updateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.doRequest("POST")
}

type PullsUpdateCmd struct {
	baseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `required:"" name:"owner"`
	Repo                string `required:"" name:"repo"`
	Number              int64  `required:"" name:"number"`
	Title               string `name:"title" help:"The title of the pull request."`
	Body                string `name:"body" help:"The contents of the pull request."`
	State               string "name:\"state\" help:\"State of this Pull Request. Either `open` or `closed`.\""
	Base                string `name:"base" help:"The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("title", c.Title)
	c.updateBody("body", c.Body)
	c.updateBody("state", c.State)
	c.updateBody("base", c.Base)
	c.updateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.doRequest("PATCH")
}

type PullsListCommitsCmd struct {
	baseCmd
	Owner   string `required:"" name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/commits"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsListFilesCmd struct {
	baseCmd
	Owner   string `required:"" name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListFilesCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/files"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsCheckIfMergedCmd struct {
	baseCmd
	Owner  string `required:"" name:"owner"`
	Repo   string `required:"" name:"repo"`
	Number int64  `required:"" name:"number"`
}

func (c *PullsCheckIfMergedCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/merge"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	return c.doRequest("GET")
}

type PullsMergeCmd struct {
	baseCmd
	Owner         string `required:"" name:"owner"`
	Repo          string `required:"" name:"repo"`
	Number        int64  `required:"" name:"number"`
	CommitTitle   string `name:"commit_title" help:"Title for the automatic commit message."`
	CommitMessage string `name:"commit_message" help:"Extra detail to append to automatic commit message."`
	Sha           string `name:"sha" help:"SHA that pull request head must match to allow merge."`
	MergeMethod   string "name:\"merge_method\" help:\"Merge method to use. Possible values are `merge`, `squash` or `rebase`. Default is `merge`.\""
}

func (c *PullsMergeCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/merge"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("commit_title", c.CommitTitle)
	c.updateBody("commit_message", c.CommitMessage)
	c.updateBody("sha", c.Sha)
	c.updateBody("merge_method", c.MergeMethod)
	return c.doRequest("PUT")
}

type PullsListReviewsCmd struct {
	baseCmd
	Owner   string `required:"" name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListReviewsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsGetReviewCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
}

func (c *PullsGetReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews/:review_id"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLPath("review_id", c.ReviewId)
	return c.doRequest("GET")
}

type PullsDeletePendingReviewCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
}

func (c *PullsDeletePendingReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews/:review_id"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLPath("review_id", c.ReviewId)
	return c.doRequest("DELETE")
}

type PullsGetCommentsForReviewCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	PerPage  int64  `name:"per_page" help:"Results per page (max 100)"`
	Page     int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsGetCommentsForReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews/:review_id/comments"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLPath("review_id", c.ReviewId)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsSubmitReviewCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	Body     string `name:"body" help:"The body text of the pull request review"`
	Event    string "required:\"\" name:\"event\" help:\"The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. When you leave this blank, the API returns _HTTP 422 (Unrecognizable entity)_ and sets the review action state to `PENDING`, which means you will need to re-submit the pull request review using a review action.\""
}

func (c *PullsSubmitReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews/:review_id/events"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLPath("review_id", c.ReviewId)
	c.updateBody("body", c.Body)
	c.updateBody("event", c.Event)
	return c.doRequest("POST")
}

type PullsDismissReviewCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	Message  string `required:"" name:"message" help:"The message for the pull request review dismissal"`
}

func (c *PullsDismissReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/reviews/:review_id/dismissals"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLPath("review_id", c.ReviewId)
	c.updateBody("message", c.Message)
	return c.doRequest("PUT")
}

type PullsListCommentsCmd struct {
	baseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	Number       int64  `required:"" name:"number"`
	Sort         string "name:\"sort\" help:\"Can be either `created` or `updated` comments.\""
	Direction    string "name:\"direction\" help:\"Can be either `asc` or `desc`. Ignored without `sort` parameter.\""
	Since        string "name:\"since\" help:\"This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Only returns comments `updated` at or after this time.\""
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/comments"
	c.updatePreview("squirrel-girl", c.SquirrelGirl)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("direction", c.Direction)
	c.updateURLQuery("since", c.Since)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsListCommentsForRepoCmd struct {
	baseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	Sort         string "name:\"sort\" help:\"Can be either `created` or `updated` comments.\""
	Direction    string "name:\"direction\" help:\"Can be either `asc` or `desc`. Ignored without `sort` parameter.\""
	Since        string "name:\"since\" help:\"This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Only returns comments `updated` at or after this time.\""
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommentsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/comments"
	c.updatePreview("squirrel-girl", c.SquirrelGirl)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("direction", c.Direction)
	c.updateURLQuery("since", c.Since)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsGetCommentCmd struct {
	baseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
}

func (c *PullsGetCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/comments/:comment_id"
	c.updatePreview("squirrel-girl", c.SquirrelGirl)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("comment_id", c.CommentId)
	return c.doRequest("GET")
}

type PullsCreateCommentCmd struct {
	baseCmd
	Owner    string `required:"" name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	Body     string `required:"" name:"body" help:"The text of the comment."`
	CommitId string "required:\"\" name:\"commit_id\" help:\"The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`.\""
	Path     string `required:"" name:"path" help:"The relative path to the file that necessitates a comment."`
	Position int64  `required:"" name:"position" help:"The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. For help finding the position value, read the note below."`
}

func (c *PullsCreateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/comments"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("body", c.Body)
	c.updateBody("commit_id", c.CommitId)
	c.updateBody("path", c.Path)
	c.updateBody("position", c.Position)
	return c.doRequest("POST")
}

type PullsCreateCommentReplyCmd struct {
	baseCmd
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	Number    int64  `required:"" name:"number"`
	Body      string `required:"" name:"body" help:"The text of the comment."`
	InReplyTo int64  `required:"" name:"in_reply_to" help:"The comment ID to reply to. **Note**: This must be the ID of a _top-level comment_, not a reply to that comment. Replies to replies are not supported."`
}

func (c *PullsCreateCommentReplyCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/comments"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("body", c.Body)
	c.updateBody("in_reply_to", c.InReplyTo)
	return c.doRequest("POST")
}

type PullsEditCommentCmd struct {
	baseCmd
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	CommentId int64  `required:"" name:"comment_id"`
	Body      string `required:"" name:"body" help:"The text of the comment."`
}

func (c *PullsEditCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/comments/:comment_id"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("comment_id", c.CommentId)
	c.updateBody("body", c.Body)
	return c.doRequest("PATCH")
}

type PullsDeleteCommentCmd struct {
	baseCmd
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	CommentId int64  `required:"" name:"comment_id"`
}

func (c *PullsDeleteCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/comments/:comment_id"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("comment_id", c.CommentId)
	return c.doRequest("DELETE")
}

type PullsListReviewRequestsCmd struct {
	baseCmd
	Owner   string `required:"" name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListReviewRequestsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/requested_reviewers"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type PullsCreateReviewRequestCmd struct {
	baseCmd
	Symmetra      bool     "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner         string   `required:"" name:"owner"`
	Repo          string   `required:"" name:"repo"`
	Number        int64    `required:"" name:"number"`
	Reviewers     []string "name:\"reviewers\" help:\"An array of user `login`s that will be requested.\""
	TeamReviewers []string "name:\"team_reviewers\" help:\"An array of team `slug`s that will be requested.\""
}

func (c *PullsCreateReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/requested_reviewers"
	c.updatePreview("symmetra", c.Symmetra)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("reviewers", c.Reviewers)
	c.updateBody("team_reviewers", c.TeamReviewers)
	return c.doRequest("POST")
}

type PullsDeleteReviewRequestCmd struct {
	baseCmd
	Owner         string   `required:"" name:"owner"`
	Repo          string   `required:"" name:"repo"`
	Number        int64    `required:"" name:"number"`
	Reviewers     []string "name:\"reviewers\" help:\"An array of user `login`s that will be removed.\""
	TeamReviewers []string "name:\"team_reviewers\" help:\"An array of team `slug`s that will be removed.\""
}

func (c *PullsDeleteReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/repos/:owner/:repo/pulls/:number/requested_reviewers"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateURLPath("number", c.Number)
	c.updateBody("reviewers", c.Reviewers)
	c.updateBody("team_reviewers", c.TeamReviewers)
	return c.doRequest("DELETE")
}
