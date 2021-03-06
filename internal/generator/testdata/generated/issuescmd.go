// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type IssuesCmd struct {
	AddAssignees             IssuesAddAssigneesCmd             `cmd:""`
	AddLabels                IssuesAddLabelsCmd                `cmd:""`
	CheckAssignee            IssuesCheckAssigneeCmd            `cmd:""`
	Create                   IssuesCreateCmd                   `cmd:""`
	CreateComment            IssuesCreateCommentCmd            `cmd:""`
	CreateLabel              IssuesCreateLabelCmd              `cmd:""`
	CreateMilestone          IssuesCreateMilestoneCmd          `cmd:""`
	DeleteComment            IssuesDeleteCommentCmd            `cmd:""`
	DeleteLabel              IssuesDeleteLabelCmd              `cmd:""`
	DeleteMilestone          IssuesDeleteMilestoneCmd          `cmd:""`
	Get                      IssuesGetCmd                      `cmd:""`
	GetComment               IssuesGetCommentCmd               `cmd:""`
	GetEvent                 IssuesGetEventCmd                 `cmd:""`
	GetLabel                 IssuesGetLabelCmd                 `cmd:""`
	GetMilestone             IssuesGetMilestoneCmd             `cmd:""`
	List                     IssuesListCmd                     `cmd:""`
	ListAssignees            IssuesListAssigneesCmd            `cmd:""`
	ListComments             IssuesListCommentsCmd             `cmd:""`
	ListCommentsForRepo      IssuesListCommentsForRepoCmd      `cmd:""`
	ListEvents               IssuesListEventsCmd               `cmd:""`
	ListEventsForRepo        IssuesListEventsForRepoCmd        `cmd:""`
	ListEventsForTimeline    IssuesListEventsForTimelineCmd    `cmd:""`
	ListForAuthenticatedUser IssuesListForAuthenticatedUserCmd `cmd:""`
	ListForOrg               IssuesListForOrgCmd               `cmd:""`
	ListForRepo              IssuesListForRepoCmd              `cmd:""`
	ListLabelsForMilestone   IssuesListLabelsForMilestoneCmd   `cmd:""`
	ListLabelsForRepo        IssuesListLabelsForRepoCmd        `cmd:""`
	ListLabelsOnIssue        IssuesListLabelsOnIssueCmd        `cmd:""`
	ListMilestonesForRepo    IssuesListMilestonesForRepoCmd    `cmd:""`
	Lock                     IssuesLockCmd                     `cmd:""`
	RemoveAllLabels          IssuesRemoveAllLabelsCmd          `cmd:""`
	RemoveAssignees          IssuesRemoveAssigneesCmd          `cmd:""`
	RemoveLabel              IssuesRemoveLabelCmd              `cmd:""`
	ReplaceAllLabels         IssuesReplaceAllLabelsCmd         `cmd:""`
	Unlock                   IssuesUnlockCmd                   `cmd:""`
	Update                   IssuesUpdateCmd                   `cmd:""`
	UpdateComment            IssuesUpdateCommentCmd            `cmd:""`
	UpdateLabel              IssuesUpdateLabelCmd              `cmd:""`
	UpdateMilestone          IssuesUpdateMilestoneCmd          `cmd:""`
}

type IssuesAddAssigneesCmd struct {
	Assignees   []string `name:"assignees"`
	IssueNumber int64    `required:"" name:"issue_number"`
	Owner       string   `name:"owner"`
	Repo        string   `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesAddAssigneesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/assignees")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("assignees", c.Assignees)
	return c.DoRequest("POST")
}

type IssuesAddLabelsCmd struct {
	IssueNumber int64    `required:"" name:"issue_number"`
	Labels      []string `required:"" name:"labels"`
	Owner       string   `name:"owner"`
	Repo        string   `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesAddLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("labels", c.Labels)
	return c.DoRequest("POST")
}

type IssuesCheckAssigneeCmd struct {
	Assignee string `required:"" name:"assignee"`
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesCheckAssigneeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/assignees/{assignee}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("assignee", c.Assignee)
	return c.DoRequest("GET")
}

type IssuesCreateCmd struct {
	Assignee  string   `name:"assignee"`
	Assignees []string `name:"assignees"`
	Body      string   `name:"body"`
	Labels    []string `name:"labels"`
	Milestone int64    `name:"milestone"`
	Owner     string   `name:"owner"`
	Repo      string   `required:"" name:"repo"`
	Title     string   `required:"" name:"title"`
	internal.BaseCmd
}

func (c *IssuesCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("assignee", c.Assignee)
	c.UpdateBody("assignees", c.Assignees)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("labels", c.Labels)
	c.UpdateBody("milestone", c.Milestone)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("POST")
}

type IssuesCreateCommentCmd struct {
	Body        string `required:"" name:"body"`
	IssueNumber int64  `required:"" name:"issue_number"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesCreateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("POST")
}

type IssuesCreateLabelCmd struct {
	Color       string `required:"" name:"color"`
	Description string `name:"description"`
	Name        string `required:"" name:"name"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesCreateLabelCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("color", c.Color)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("name", c.Name)
	return c.DoRequest("POST")
}

type IssuesCreateMilestoneCmd struct {
	Description string `name:"description"`
	DueOn       string `name:"due_on"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	State       string `name:"state"`
	Title       string `required:"" name:"title"`
	internal.BaseCmd
}

func (c *IssuesCreateMilestoneCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("due_on", c.DueOn)
	c.UpdateBody("state", c.State)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("POST")
}

type IssuesDeleteCommentCmd struct {
	CommentId int64  `required:"" name:"comment_id"`
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesDeleteCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/comments/{comment_id}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("DELETE")
}

type IssuesDeleteLabelCmd struct {
	Name  string `required:"" name:"name"`
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesDeleteLabelCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/labels/{name}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("name", c.Name)
	return c.DoRequest("DELETE")
}

type IssuesDeleteMilestoneCmd struct {
	MilestoneNumber int64  `required:"" name:"milestone_number"`
	Owner           string `name:"owner"`
	Repo            string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesDeleteMilestoneCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones/{milestone_number}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("milestone_number", c.MilestoneNumber)
	return c.DoRequest("DELETE")
}

type IssuesGetCmd struct {
	IssueNumber  int64  `required:"" name:"issue_number"`
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *IssuesGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesGetCommentCmd struct {
	CommentId    int64  `required:"" name:"comment_id"`
	MachineMan   bool   `name:"machine-man-preview"`
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *IssuesGetCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/comments/{comment_id}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesGetEventCmd struct {
	EventId    int64  `required:"" name:"event_id"`
	MachineMan bool   `name:"machine-man-preview"`
	Owner      string `name:"owner"`
	Repo       string `required:"" name:"repo"`
	SailorV    bool   `name:"sailor-v-preview"`
	Starfox    bool   `name:"starfox-preview"`
	internal.BaseCmd
}

func (c *IssuesGetEventCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/events/{event_id}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("event_id", c.EventId)
	c.UpdatePreview("starfox", c.Starfox)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("sailor-v", c.SailorV)
	return c.DoRequest("GET")
}

type IssuesGetLabelCmd struct {
	Name  string `required:"" name:"name"`
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesGetLabelCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/labels/{name}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("name", c.Name)
	return c.DoRequest("GET")
}

type IssuesGetMilestoneCmd struct {
	MilestoneNumber int64  `required:"" name:"milestone_number"`
	Owner           string `name:"owner"`
	Repo            string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesGetMilestoneCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones/{milestone_number}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("milestone_number", c.MilestoneNumber)
	return c.DoRequest("GET")
}

type IssuesListAssigneesCmd struct {
	Owner   string `name:"owner"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Repo    string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesListAssigneesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/assignees")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type IssuesListCmd struct {
	Direction    string `name:"direction"`
	Filter       string `name:"filter"`
	Labels       string `name:"labels"`
	MachineMan   bool   `name:"machine-man-preview"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	State        string `name:"state"`
	internal.BaseCmd
}

func (c *IssuesListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/issues")
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("labels", c.Labels)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListCommentsCmd struct {
	IssueNumber  int64  `required:"" name:"issue_number"`
	Owner        string `name:"owner"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Repo         string `required:"" name:"repo"`
	Since        string `name:"since"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *IssuesListCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListCommentsForRepoCmd struct {
	Direction    string `name:"direction"`
	Owner        string `name:"owner"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Repo         string `required:"" name:"repo"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *IssuesListCommentsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListEventsCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Owner       string `name:"owner"`
	Page        int64  `name:"page"`
	PerPage     int64  `name:"per_page"`
	Repo        string `required:"" name:"repo"`
	SailorV     bool   `name:"sailor-v-preview"`
	Starfox     bool   `name:"starfox-preview"`
	internal.BaseCmd
}

func (c *IssuesListEventsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/events")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("starfox", c.Starfox)
	c.UpdatePreview("sailor-v", c.SailorV)
	return c.DoRequest("GET")
}

type IssuesListEventsForRepoCmd struct {
	Owner   string `name:"owner"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Repo    string `required:"" name:"repo"`
	SailorV bool   `name:"sailor-v-preview"`
	Starfox bool   `name:"starfox-preview"`
	internal.BaseCmd
}

func (c *IssuesListEventsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/events")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("starfox", c.Starfox)
	c.UpdatePreview("sailor-v", c.SailorV)
	return c.DoRequest("GET")
}

type IssuesListEventsForTimelineCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Mockingbird bool   `required:"" name:"mockingbird-preview"`
	Owner       string `name:"owner"`
	Page        int64  `name:"page"`
	PerPage     int64  `name:"per_page"`
	Repo        string `required:"" name:"repo"`
	Starfox     bool   `name:"starfox-preview"`
	internal.BaseCmd
}

func (c *IssuesListEventsForTimelineCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/timeline")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("mockingbird", c.Mockingbird)
	c.UpdatePreview("starfox", c.Starfox)
	return c.DoRequest("GET")
}

type IssuesListForAuthenticatedUserCmd struct {
	Direction    string `name:"direction"`
	Filter       string `name:"filter"`
	Labels       string `name:"labels"`
	MachineMan   bool   `name:"machine-man-preview"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	State        string `name:"state"`
	internal.BaseCmd
}

func (c *IssuesListForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/issues")
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("labels", c.Labels)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListForOrgCmd struct {
	Direction    string `name:"direction"`
	Filter       string `name:"filter"`
	Labels       string `name:"labels"`
	MachineMan   bool   `name:"machine-man-preview"`
	Org          string `required:"" name:"org"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	State        string `name:"state"`
	internal.BaseCmd
}

func (c *IssuesListForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/issues")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("labels", c.Labels)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListForRepoCmd struct {
	Assignee     string `name:"assignee"`
	Creator      string `name:"creator"`
	Direction    string `name:"direction"`
	Labels       string `name:"labels"`
	MachineMan   bool   `name:"machine-man-preview"`
	Mentioned    string `name:"mentioned"`
	Milestone    string `name:"milestone"`
	Owner        string `name:"owner"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Repo         string `required:"" name:"repo"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	State        string `name:"state"`
	internal.BaseCmd
}

func (c *IssuesListForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("milestone", c.Milestone)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("assignee", c.Assignee)
	c.UpdateURLQuery("creator", c.Creator)
	c.UpdateURLQuery("mentioned", c.Mentioned)
	c.UpdateURLQuery("labels", c.Labels)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type IssuesListLabelsForMilestoneCmd struct {
	MilestoneNumber int64  `required:"" name:"milestone_number"`
	Owner           string `name:"owner"`
	Page            int64  `name:"page"`
	PerPage         int64  `name:"per_page"`
	Repo            string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesListLabelsForMilestoneCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones/{milestone_number}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("milestone_number", c.MilestoneNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type IssuesListLabelsForRepoCmd struct {
	Owner   string `name:"owner"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Repo    string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesListLabelsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type IssuesListLabelsOnIssueCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Owner       string `name:"owner"`
	Page        int64  `name:"page"`
	PerPage     int64  `name:"per_page"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesListLabelsOnIssueCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type IssuesListMilestonesForRepoCmd struct {
	Direction string `name:"direction"`
	Owner     string `name:"owner"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Repo      string `required:"" name:"repo"`
	Sort      string `name:"sort"`
	State     string `name:"state"`
	internal.BaseCmd
}

func (c *IssuesListMilestonesForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type IssuesLockCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	LockReason  string `name:"lock_reason"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	SailorV     bool   `name:"sailor-v-preview"`
	internal.BaseCmd
}

func (c *IssuesLockCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/lock")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdatePreview("sailor-v", c.SailorV)
	c.UpdateBody("lock_reason", c.LockReason)
	return c.DoRequest("PUT")
}

type IssuesRemoveAllLabelsCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesRemoveAllLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	return c.DoRequest("DELETE")
}

type IssuesRemoveAssigneesCmd struct {
	Assignees   []string `name:"assignees"`
	IssueNumber int64    `required:"" name:"issue_number"`
	Owner       string   `name:"owner"`
	Repo        string   `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesRemoveAssigneesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/assignees")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("assignees", c.Assignees)
	return c.DoRequest("DELETE")
}

type IssuesRemoveLabelCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Name        string `required:"" name:"name"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesRemoveLabelCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/labels/{name}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateURLPath("name", c.Name)
	return c.DoRequest("DELETE")
}

type IssuesReplaceAllLabelsCmd struct {
	IssueNumber int64    `required:"" name:"issue_number"`
	Labels      []string `name:"labels"`
	Owner       string   `name:"owner"`
	Repo        string   `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesReplaceAllLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/labels")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("labels", c.Labels)
	return c.DoRequest("PUT")
}

type IssuesUnlockCmd struct {
	IssueNumber int64  `required:"" name:"issue_number"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesUnlockCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}/lock")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	return c.DoRequest("DELETE")
}

type IssuesUpdateCmd struct {
	Assignee    string   `name:"assignee"`
	Assignees   []string `name:"assignees"`
	Body        string   `name:"body"`
	IssueNumber int64    `required:"" name:"issue_number"`
	Labels      []string `name:"labels"`
	Milestone   int64    `name:"milestone"`
	Owner       string   `name:"owner"`
	Repo        string   `required:"" name:"repo"`
	State       string   `name:"state"`
	Title       string   `name:"title"`
	internal.BaseCmd
}

func (c *IssuesUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/{issue_number}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("issue_number", c.IssueNumber)
	c.UpdateBody("assignee", c.Assignee)
	c.UpdateBody("assignees", c.Assignees)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("labels", c.Labels)
	c.UpdateBody("milestone", c.Milestone)
	c.UpdateBody("state", c.State)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("PATCH")
}

type IssuesUpdateCommentCmd struct {
	Body      string `required:"" name:"body"`
	CommentId int64  `required:"" name:"comment_id"`
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesUpdateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/issues/comments/{comment_id}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PATCH")
}

type IssuesUpdateLabelCmd struct {
	Color       string `name:"color"`
	Description string `name:"description"`
	Name        string `required:"" name:"name"`
	NewName     string `name:"new_name"`
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *IssuesUpdateLabelCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/labels/{name}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("name", c.Name)
	c.UpdateBody("color", c.Color)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("new_name", c.NewName)
	return c.DoRequest("PATCH")
}

type IssuesUpdateMilestoneCmd struct {
	Description     string `name:"description"`
	DueOn           string `name:"due_on"`
	MilestoneNumber int64  `required:"" name:"milestone_number"`
	Owner           string `name:"owner"`
	Repo            string `required:"" name:"repo"`
	State           string `name:"state"`
	Title           string `name:"title"`
	internal.BaseCmd
}

func (c *IssuesUpdateMilestoneCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/milestones/{milestone_number}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("milestone_number", c.MilestoneNumber)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("due_on", c.DueOn)
	c.UpdateBody("state", c.State)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("PATCH")
}
