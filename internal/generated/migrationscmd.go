// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type MigrationsCmd struct {
	CancelImport                      MigrationsCancelImportCmd                      `cmd:""`
	DeleteArchiveForAuthenticatedUser MigrationsDeleteArchiveForAuthenticatedUserCmd `cmd:""`
	DeleteArchiveForOrg               MigrationsDeleteArchiveForOrgCmd               `cmd:""`
	DownloadArchiveForOrg             MigrationsDownloadArchiveForOrgCmd             `cmd:""`
	GetArchiveForAuthenticatedUser    MigrationsGetArchiveForAuthenticatedUserCmd    `cmd:""`
	GetCommitAuthors                  MigrationsGetCommitAuthorsCmd                  `cmd:""`
	GetImportProgress                 MigrationsGetImportProgressCmd                 `cmd:""`
	GetLargeFiles                     MigrationsGetLargeFilesCmd                     `cmd:""`
	GetStatusForAuthenticatedUser     MigrationsGetStatusForAuthenticatedUserCmd     `cmd:""`
	GetStatusForOrg                   MigrationsGetStatusForOrgCmd                   `cmd:""`
	ListForAuthenticatedUser          MigrationsListForAuthenticatedUserCmd          `cmd:""`
	ListForOrg                        MigrationsListForOrgCmd                        `cmd:""`
	ListReposForOrg                   MigrationsListReposForOrgCmd                   `cmd:""`
	ListReposForUser                  MigrationsListReposForUserCmd                  `cmd:""`
	MapCommitAuthor                   MigrationsMapCommitAuthorCmd                   `cmd:""`
	SetLfsPreference                  MigrationsSetLfsPreferenceCmd                  `cmd:""`
	StartForAuthenticatedUser         MigrationsStartForAuthenticatedUserCmd         `cmd:""`
	StartForOrg                       MigrationsStartForOrgCmd                       `cmd:""`
	StartImport                       MigrationsStartImportCmd                       `cmd:""`
	UnlockRepoForAuthenticatedUser    MigrationsUnlockRepoForAuthenticatedUserCmd    `cmd:""`
	UnlockRepoForOrg                  MigrationsUnlockRepoForOrgCmd                  `cmd:""`
	UpdateImport                      MigrationsUpdateImportCmd                      `cmd:""`
}

type MigrationsCancelImportCmd struct {
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *MigrationsCancelImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("DELETE")
}

type MigrationsDeleteArchiveForAuthenticatedUserCmd struct {
	MigrationId int64 `required:"" name:"migration_id"`
	Wyandotte   bool  "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsDeleteArchiveForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsDeleteArchiveForOrgCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	Org         string `required:"" name:"org"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsDeleteArchiveForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsDownloadArchiveForOrgCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	Org         string `required:"" name:"org"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsDownloadArchiveForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetArchiveForAuthenticatedUserCmd struct {
	MigrationId int64 `required:"" name:"migration_id"`
	Wyandotte   bool  "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsGetArchiveForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetCommitAuthorsCmd struct {
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	Since string `name:"since"`
	internal.BaseCmd
}

func (c *MigrationsGetCommitAuthorsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import/authors")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("since", c.Since)
	return c.DoRequest("GET")
}

type MigrationsGetImportProgressCmd struct {
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *MigrationsGetImportProgressCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type MigrationsGetLargeFilesCmd struct {
	Owner string `name:"owner"`
	Repo  string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *MigrationsGetLargeFilesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import/large_files")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type MigrationsGetStatusForAuthenticatedUserCmd struct {
	MigrationId int64 `required:"" name:"migration_id"`
	Wyandotte   bool  "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsGetStatusForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetStatusForOrgCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	Org         string `required:"" name:"org"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsGetStatusForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListForAuthenticatedUserCmd struct {
	Page      int64 `name:"page"`
	PerPage   int64 `name:"per_page"`
	Wyandotte bool  "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsListForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListForOrgCmd struct {
	Org       string `required:"" name:"org"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Wyandotte bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsListForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListReposForOrgCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	Org         string `required:"" name:"org"`
	Page        int64  `name:"page"`
	PerPage     int64  `name:"per_page"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsListReposForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/repositories")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListReposForUserCmd struct {
	MigrationId int64 `required:"" name:"migration_id"`
	Page        int64 `name:"page"`
	PerPage     int64 `name:"per_page"`
	Wyandotte   bool  "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsListReposForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/{migration_id}/repositories")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsMapCommitAuthorCmd struct {
	AuthorId int64  `required:"" name:"author_id"`
	Email    string `name:"email"`
	Name     string `name:"name"`
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *MigrationsMapCommitAuthorCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import/authors/{author_id}")
	c.UpdateURLPath("author_id", c.AuthorId)
	c.UpdateBody("email", c.Email)
	c.UpdateBody("name", c.Name)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("PATCH")
}

type MigrationsSetLfsPreferenceCmd struct {
	Owner  string `name:"owner"`
	Repo   string `required:"" name:"repo"`
	UseLfs string `required:"" name:"use_lfs"`
	internal.BaseCmd
}

func (c *MigrationsSetLfsPreferenceCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import/lfs")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("use_lfs", c.UseLfs)
	return c.DoRequest("PATCH")
}

type MigrationsStartForAuthenticatedUserCmd struct {
	ExcludeAttachments bool     `name:"exclude_attachments"`
	LockRepositories   bool     `name:"lock_repositories"`
	Repositories       []string `required:"" name:"repositories"`
	internal.BaseCmd
}

func (c *MigrationsStartForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations")
	c.UpdateBody("exclude_attachments", c.ExcludeAttachments)
	c.UpdateBody("lock_repositories", c.LockRepositories)
	c.UpdateBody("repositories", c.Repositories)
	return c.DoRequest("POST")
}

type MigrationsStartForOrgCmd struct {
	ExcludeAttachments bool     `name:"exclude_attachments"`
	LockRepositories   bool     `name:"lock_repositories"`
	Org                string   `required:"" name:"org"`
	Repositories       []string `required:"" name:"repositories"`
	internal.BaseCmd
}

func (c *MigrationsStartForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations")
	c.UpdateBody("exclude_attachments", c.ExcludeAttachments)
	c.UpdateBody("lock_repositories", c.LockRepositories)
	c.UpdateURLPath("org", c.Org)
	c.UpdateBody("repositories", c.Repositories)
	return c.DoRequest("POST")
}

type MigrationsStartImportCmd struct {
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	TfvcProject string `name:"tfvc_project"`
	Vcs         string `name:"vcs"`
	VcsPassword string `name:"vcs_password"`
	VcsUrl      string `required:"" name:"vcs_url"`
	VcsUsername string `name:"vcs_username"`
	internal.BaseCmd
}

func (c *MigrationsStartImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("tfvc_project", c.TfvcProject)
	c.UpdateBody("vcs", c.Vcs)
	c.UpdateBody("vcs_password", c.VcsPassword)
	c.UpdateBody("vcs_url", c.VcsUrl)
	c.UpdateBody("vcs_username", c.VcsUsername)
	return c.DoRequest("PUT")
}

type MigrationsUnlockRepoForAuthenticatedUserCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	RepoName    string `required:"" name:"repo_name"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsUnlockRepoForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/repos/{repo_name}/lock")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("repo_name", c.RepoName)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsUnlockRepoForOrgCmd struct {
	MigrationId int64  `required:"" name:"migration_id"`
	Org         string `required:"" name:"org"`
	RepoName    string `required:"" name:"repo_name"`
	Wyandotte   bool   "name:\"wyandotte-preview\" required:\"\" help:\"To access the Migrations API, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.wyandotte-preview+json\n```\""
	internal.BaseCmd
}

func (c *MigrationsUnlockRepoForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("repo_name", c.RepoName)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsUpdateImportCmd struct {
	Owner       string `name:"owner"`
	Repo        string `required:"" name:"repo"`
	VcsPassword string `name:"vcs_password"`
	VcsUsername string `name:"vcs_username"`
	internal.BaseCmd
}

func (c *MigrationsUpdateImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/import")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("vcs_password", c.VcsPassword)
	c.UpdateBody("vcs_username", c.VcsUsername)
	return c.DoRequest("PATCH")
}
