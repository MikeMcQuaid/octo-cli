// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type AppsCmd struct {
	AddRepoToInstallation                        AppsAddRepoToInstallationCmd                        `cmd:""`
	CheckAuthorization                           AppsCheckAuthorizationCmd                           `cmd:""`
	CheckToken                                   AppsCheckTokenCmd                                   `cmd:""`
	CreateContentAttachment                      AppsCreateContentAttachmentCmd                      `cmd:""`
	CreateFromManifest                           AppsCreateFromManifestCmd                           `cmd:""`
	CreateInstallationToken                      AppsCreateInstallationTokenCmd                      `cmd:""`
	DeleteAuthorization                          AppsDeleteAuthorizationCmd                          `cmd:""`
	DeleteInstallation                           AppsDeleteInstallationCmd                           `cmd:""`
	DeleteToken                                  AppsDeleteTokenCmd                                  `cmd:""`
	GetAuthenticated                             AppsGetAuthenticatedCmd                             `cmd:""`
	GetBySlug                                    AppsGetBySlugCmd                                    `cmd:""`
	GetInstallation                              AppsGetInstallationCmd                              `cmd:""`
	GetOrgInstallation                           AppsGetOrgInstallationCmd                           `cmd:""`
	GetRepoInstallation                          AppsGetRepoInstallationCmd                          `cmd:""`
	GetSubscriptionPlanForAccount                AppsGetSubscriptionPlanForAccountCmd                `cmd:""`
	GetSubscriptionPlanForAccountStubbed         AppsGetSubscriptionPlanForAccountStubbedCmd         `cmd:""`
	GetUserInstallation                          AppsGetUserInstallationCmd                          `cmd:""`
	ListAccountsForPlan                          AppsListAccountsForPlanCmd                          `cmd:""`
	ListAccountsForPlanStubbed                   AppsListAccountsForPlanStubbedCmd                   `cmd:""`
	ListInstallationReposForAuthenticatedUser    AppsListInstallationReposForAuthenticatedUserCmd    `cmd:""`
	ListInstallations                            AppsListInstallationsCmd                            `cmd:""`
	ListInstallationsForAuthenticatedUser        AppsListInstallationsForAuthenticatedUserCmd        `cmd:""`
	ListPlans                                    AppsListPlansCmd                                    `cmd:""`
	ListPlansStubbed                             AppsListPlansStubbedCmd                             `cmd:""`
	ListRepos                                    AppsListReposCmd                                    `cmd:""`
	ListSubscriptionsForAuthenticatedUser        AppsListSubscriptionsForAuthenticatedUserCmd        `cmd:""`
	ListSubscriptionsForAuthenticatedUserStubbed AppsListSubscriptionsForAuthenticatedUserStubbedCmd `cmd:""`
	RemoveRepoFromInstallation                   AppsRemoveRepoFromInstallationCmd                   `cmd:""`
	ResetAuthorization                           AppsResetAuthorizationCmd                           `cmd:""`
	ResetToken                                   AppsResetTokenCmd                                   `cmd:""`
	RevokeAuthorizationForApplication            AppsRevokeAuthorizationForApplicationCmd            `cmd:""`
	RevokeGrantForApplication                    AppsRevokeGrantForApplicationCmd                    `cmd:""`
	RevokeInstallationToken                      AppsRevokeInstallationTokenCmd                      `cmd:""`
}

type AppsAddRepoToInstallationCmd struct {
	InstallationId int64 `required:"" name:"installation_id"`
	MachineMan     bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	RepositoryId   int64 `required:"" name:"repository_id"`
	internal.BaseCmd
}

func (c *AppsAddRepoToInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/installations/{installation_id}/repositories/{repository_id}")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("repository_id", c.RepositoryId)
	return c.DoRequest("PUT")
}

type AppsCheckAuthorizationCmd struct {
	AccessToken string `required:"" name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsCheckAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/tokens/{access_token}")
	c.UpdateURLPath("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("GET")
}

type AppsCheckTokenCmd struct {
	AccessToken string `name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsCheckTokenCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/token")
	c.UpdateBody("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("POST")
}

type AppsCreateContentAttachmentCmd struct {
	Body               string `required:"" name:"body"`
	ContentReferenceId int64  `required:"" name:"content_reference_id"`
	Corsair            bool   "name:\"corsair-preview\" required:\"\" help:\"To access the Content Attachments API during the preview period, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.corsair-preview+json\n```\""
	Title              string `required:"" name:"title"`
	internal.BaseCmd
}

func (c *AppsCreateContentAttachmentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/content_references/{content_reference_id}/attachments")
	c.UpdateBody("body", c.Body)
	c.UpdateURLPath("content_reference_id", c.ContentReferenceId)
	c.UpdatePreview("corsair", c.Corsair)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("POST")
}

type AppsCreateFromManifestCmd struct {
	Code string `required:"" name:"code"`
	internal.BaseCmd
}

func (c *AppsCreateFromManifestCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app-manifests/{code}/conversions")
	c.UpdateURLPath("code", c.Code)
	return c.DoRequest("POST")
}

type AppsCreateInstallationTokenCmd struct {
	InstallationId int64   `required:"" name:"installation_id"`
	MachineMan     bool    "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	RepositoryIds  []int64 `name:"repository_ids"`
	internal.BaseCmd
}

func (c *AppsCreateInstallationTokenCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app/installations/{installation_id}/access_tokens")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateBody("repository_ids", c.RepositoryIds)
	return c.DoRequest("POST")
}

type AppsDeleteAuthorizationCmd struct {
	AccessToken string `name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsDeleteAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/grant")
	c.UpdateBody("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("DELETE")
}

type AppsDeleteInstallationCmd struct {
	InstallationId int64 `required:"" name:"installation_id"`
	MachineMan     bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	internal.BaseCmd
}

func (c *AppsDeleteInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app/installations/{installation_id}")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	return c.DoRequest("DELETE")
}

type AppsDeleteTokenCmd struct {
	AccessToken string `name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsDeleteTokenCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/token")
	c.UpdateBody("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("DELETE")
}

type AppsGetAuthenticatedCmd struct {
	MachineMan bool "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	internal.BaseCmd
}

func (c *AppsGetAuthenticatedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app")
	c.UpdatePreview("machine-man", c.MachineMan)
	return c.DoRequest("GET")
}

type AppsGetBySlugCmd struct {
	AppSlug    string `required:"" name:"app_slug"`
	MachineMan bool   "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	internal.BaseCmd
}

func (c *AppsGetBySlugCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/apps/{app_slug}")
	c.UpdateURLPath("app_slug", c.AppSlug)
	c.UpdatePreview("machine-man", c.MachineMan)
	return c.DoRequest("GET")
}

type AppsGetInstallationCmd struct {
	InstallationId int64 `required:"" name:"installation_id"`
	MachineMan     bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	internal.BaseCmd
}

func (c *AppsGetInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app/installations/{installation_id}")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	return c.DoRequest("GET")
}

type AppsGetOrgInstallationCmd struct {
	MachineMan bool   "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Org        string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *AppsGetOrgInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/installation")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("GET")
}

type AppsGetRepoInstallationCmd struct {
	MachineMan bool   "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Owner      string `name:"owner"`
	Repo       string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *AppsGetRepoInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/installation")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type AppsGetSubscriptionPlanForAccountCmd struct {
	AccountId int64 `required:"" name:"account_id"`
	internal.BaseCmd
}

func (c *AppsGetSubscriptionPlanForAccountCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/accounts/{account_id}")
	c.UpdateURLPath("account_id", c.AccountId)
	return c.DoRequest("GET")
}

type AppsGetSubscriptionPlanForAccountStubbedCmd struct {
	AccountId int64 `required:"" name:"account_id"`
	internal.BaseCmd
}

func (c *AppsGetSubscriptionPlanForAccountStubbedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/stubbed/accounts/{account_id}")
	c.UpdateURLPath("account_id", c.AccountId)
	return c.DoRequest("GET")
}

type AppsGetUserInstallationCmd struct {
	MachineMan bool   "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Username   string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *AppsGetUserInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/installation")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type AppsListAccountsForPlanCmd struct {
	Direction string `name:"direction"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	PlanId    int64  `required:"" name:"plan_id"`
	Sort      string `name:"sort"`
	internal.BaseCmd
}

func (c *AppsListAccountsForPlanCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/plans/{plan_id}/accounts")
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("plan_id", c.PlanId)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type AppsListAccountsForPlanStubbedCmd struct {
	Direction string `name:"direction"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	PlanId    int64  `required:"" name:"plan_id"`
	Sort      string `name:"sort"`
	internal.BaseCmd
}

func (c *AppsListAccountsForPlanStubbedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/stubbed/plans/{plan_id}/accounts")
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("plan_id", c.PlanId)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type AppsListInstallationReposForAuthenticatedUserCmd struct {
	InstallationId int64 `required:"" name:"installation_id"`
	MachineMan     bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Mercy          bool  "name:\"mercy-preview\" help:\"The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.mercy-preview+json\n```\""
	Page           int64 `name:"page"`
	PerPage        int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListInstallationReposForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/installations/{installation_id}/repositories")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("mercy", c.Mercy)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListInstallationsCmd struct {
	MachineMan bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Page       int64 `name:"page"`
	PerPage    int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListInstallationsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/app/installations")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListInstallationsForAuthenticatedUserCmd struct {
	MachineMan bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Page       int64 `name:"page"`
	PerPage    int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListInstallationsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/installations")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListPlansCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListPlansCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/plans")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListPlansStubbedCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListPlansStubbedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/marketplace_listing/stubbed/plans")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListReposCmd struct {
	MachineMan bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Mercy      bool  "name:\"mercy-preview\" help:\"The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.mercy-preview+json\n```\""
	Page       int64 `name:"page"`
	PerPage    int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListReposCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/installation/repositories")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdatePreview("mercy", c.Mercy)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListSubscriptionsForAuthenticatedUserCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListSubscriptionsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/marketplace_purchases")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsListSubscriptionsForAuthenticatedUserStubbedCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *AppsListSubscriptionsForAuthenticatedUserStubbedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/marketplace_purchases/stubbed")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type AppsRemoveRepoFromInstallationCmd struct {
	InstallationId int64 `required:"" name:"installation_id"`
	MachineMan     bool  "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	RepositoryId   int64 `required:"" name:"repository_id"`
	internal.BaseCmd
}

func (c *AppsRemoveRepoFromInstallationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/installations/{installation_id}/repositories/{repository_id}")
	c.UpdateURLPath("installation_id", c.InstallationId)
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("repository_id", c.RepositoryId)
	return c.DoRequest("DELETE")
}

type AppsResetAuthorizationCmd struct {
	AccessToken string `required:"" name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsResetAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/tokens/{access_token}")
	c.UpdateURLPath("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("POST")
}

type AppsResetTokenCmd struct {
	AccessToken string `name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsResetTokenCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/token")
	c.UpdateBody("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("PATCH")
}

type AppsRevokeAuthorizationForApplicationCmd struct {
	AccessToken string `required:"" name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsRevokeAuthorizationForApplicationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/tokens/{access_token}")
	c.UpdateURLPath("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("DELETE")
}

type AppsRevokeGrantForApplicationCmd struct {
	AccessToken string `required:"" name:"access_token"`
	ClientId    string `required:"" name:"client_id"`
	internal.BaseCmd
}

func (c *AppsRevokeGrantForApplicationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/applications/{client_id}/grants/{access_token}")
	c.UpdateURLPath("access_token", c.AccessToken)
	c.UpdateURLPath("client_id", c.ClientId)
	return c.DoRequest("DELETE")
}

type AppsRevokeInstallationTokenCmd struct {
	internal.BaseCmd
}

func (c *AppsRevokeInstallationTokenCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/installation/token")
	return c.DoRequest("DELETE")
}
