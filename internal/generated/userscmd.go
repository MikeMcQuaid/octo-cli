// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type UsersCmd struct {
	AddEmails                         UsersAddEmailsCmd                         `cmd:""`
	Block                             UsersBlockCmd                             `cmd:""`
	CheckBlocked                      UsersCheckBlockedCmd                      `cmd:""`
	CheckFollowing                    UsersCheckFollowingCmd                    `cmd:""`
	CheckFollowingForUser             UsersCheckFollowingForUserCmd             `cmd:""`
	CreateGpgKey                      UsersCreateGpgKeyCmd                      `cmd:""`
	CreatePublicKey                   UsersCreatePublicKeyCmd                   `cmd:""`
	DeleteEmails                      UsersDeleteEmailsCmd                      `cmd:""`
	DeleteGpgKey                      UsersDeleteGpgKeyCmd                      `cmd:""`
	DeletePublicKey                   UsersDeletePublicKeyCmd                   `cmd:""`
	Follow                            UsersFollowCmd                            `cmd:""`
	GetAuthenticated                  UsersGetAuthenticatedCmd                  `cmd:""`
	GetByUsername                     UsersGetByUsernameCmd                     `cmd:""`
	GetContextForUser                 UsersGetContextForUserCmd                 `cmd:""`
	GetGpgKey                         UsersGetGpgKeyCmd                         `cmd:""`
	GetPublicKey                      UsersGetPublicKeyCmd                      `cmd:""`
	List                              UsersListCmd                              `cmd:""`
	ListBlocked                       UsersListBlockedCmd                       `cmd:""`
	ListEmails                        UsersListEmailsCmd                        `cmd:""`
	ListFollowedByAuthenticated       UsersListFollowedByAuthenticatedCmd       `cmd:""`
	ListFollowersForAuthenticatedUser UsersListFollowersForAuthenticatedUserCmd `cmd:""`
	ListFollowersForUser              UsersListFollowersForUserCmd              `cmd:""`
	ListFollowingForUser              UsersListFollowingForUserCmd              `cmd:""`
	ListGpgKeys                       UsersListGpgKeysCmd                       `cmd:""`
	ListGpgKeysForUser                UsersListGpgKeysForUserCmd                `cmd:""`
	ListPublicEmails                  UsersListPublicEmailsCmd                  `cmd:""`
	ListPublicKeys                    UsersListPublicKeysCmd                    `cmd:""`
	ListPublicKeysForUser             UsersListPublicKeysForUserCmd             `cmd:""`
	TogglePrimaryEmailVisibility      UsersTogglePrimaryEmailVisibilityCmd      `cmd:""`
	Unblock                           UsersUnblockCmd                           `cmd:""`
	Unfollow                          UsersUnfollowCmd                          `cmd:""`
	UpdateAuthenticated               UsersUpdateAuthenticatedCmd               `cmd:""`
}

type UsersAddEmailsCmd struct {
	Emails []string `required:"" name:"emails"`
	internal.BaseCmd
}

func (c *UsersAddEmailsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/emails")
	c.UpdateBody("emails", c.Emails)
	return c.DoRequest("POST")
}

type UsersBlockCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersBlockCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/blocks/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type UsersCheckBlockedCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersCheckBlockedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/blocks/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersCheckFollowingCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersCheckFollowingCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/following/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersCheckFollowingForUserCmd struct {
	TargetUser string `required:"" name:"target_user"`
	Username   string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersCheckFollowingForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/following/{target_user}")
	c.UpdateURLPath("target_user", c.TargetUser)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersCreateGpgKeyCmd struct {
	ArmoredPublicKey string `name:"armored_public_key"`
	internal.BaseCmd
}

func (c *UsersCreateGpgKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/gpg_keys")
	c.UpdateBody("armored_public_key", c.ArmoredPublicKey)
	return c.DoRequest("POST")
}

type UsersCreatePublicKeyCmd struct {
	Key   string `name:"key"`
	Title string `name:"title"`
	internal.BaseCmd
}

func (c *UsersCreatePublicKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/keys")
	c.UpdateBody("key", c.Key)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("POST")
}

type UsersDeleteEmailsCmd struct {
	Emails []string `required:"" name:"emails"`
	internal.BaseCmd
}

func (c *UsersDeleteEmailsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/emails")
	c.UpdateBody("emails", c.Emails)
	return c.DoRequest("DELETE")
}

type UsersDeleteGpgKeyCmd struct {
	GpgKeyId int64 `required:"" name:"gpg_key_id"`
	internal.BaseCmd
}

func (c *UsersDeleteGpgKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/gpg_keys/{gpg_key_id}")
	c.UpdateURLPath("gpg_key_id", c.GpgKeyId)
	return c.DoRequest("DELETE")
}

type UsersDeletePublicKeyCmd struct {
	KeyId int64 `required:"" name:"key_id"`
	internal.BaseCmd
}

func (c *UsersDeletePublicKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/keys/{key_id}")
	c.UpdateURLPath("key_id", c.KeyId)
	return c.DoRequest("DELETE")
}

type UsersFollowCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersFollowCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/following/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type UsersGetAuthenticatedCmd struct {
	internal.BaseCmd
}

func (c *UsersGetAuthenticatedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user")
	return c.DoRequest("GET")
}

type UsersGetByUsernameCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersGetByUsernameCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersGetContextForUserCmd struct {
	SubjectId   string `name:"subject_id"`
	SubjectType string `name:"subject_type"`
	Username    string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersGetContextForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/hovercard")
	c.UpdateURLQuery("subject_id", c.SubjectId)
	c.UpdateURLQuery("subject_type", c.SubjectType)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersGetGpgKeyCmd struct {
	GpgKeyId int64 `required:"" name:"gpg_key_id"`
	internal.BaseCmd
}

func (c *UsersGetGpgKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/gpg_keys/{gpg_key_id}")
	c.UpdateURLPath("gpg_key_id", c.GpgKeyId)
	return c.DoRequest("GET")
}

type UsersGetPublicKeyCmd struct {
	KeyId int64 `required:"" name:"key_id"`
	internal.BaseCmd
}

func (c *UsersGetPublicKeyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/keys/{key_id}")
	c.UpdateURLPath("key_id", c.KeyId)
	return c.DoRequest("GET")
}

type UsersListBlockedCmd struct {
	internal.BaseCmd
}

func (c *UsersListBlockedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/blocks")
	return c.DoRequest("GET")
}

type UsersListCmd struct {
	Since string `name:"since"`
	internal.BaseCmd
}

func (c *UsersListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users")
	c.UpdateURLQuery("since", c.Since)
	return c.DoRequest("GET")
}

type UsersListEmailsCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListEmailsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/emails")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListFollowedByAuthenticatedCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListFollowedByAuthenticatedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/following")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListFollowersForAuthenticatedUserCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListFollowersForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/followers")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListFollowersForUserCmd struct {
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersListFollowersForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/followers")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersListFollowingForUserCmd struct {
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersListFollowingForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/following")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersListGpgKeysCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListGpgKeysCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/gpg_keys")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListGpgKeysForUserCmd struct {
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersListGpgKeysForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/gpg_keys")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersListPublicEmailsCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListPublicEmailsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/public_emails")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListPublicKeysCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *UsersListPublicKeysCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/keys")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type UsersListPublicKeysForUserCmd struct {
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersListPublicKeysForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/keys")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type UsersTogglePrimaryEmailVisibilityCmd struct {
	Email      string `required:"" name:"email"`
	Visibility string `required:"" name:"visibility"`
	internal.BaseCmd
}

func (c *UsersTogglePrimaryEmailVisibilityCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/email/visibility")
	c.UpdateBody("email", c.Email)
	c.UpdateBody("visibility", c.Visibility)
	return c.DoRequest("PATCH")
}

type UsersUnblockCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersUnblockCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/blocks/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type UsersUnfollowCmd struct {
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *UsersUnfollowCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/following/{username}")
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type UsersUpdateAuthenticatedCmd struct {
	Bio      string `name:"bio"`
	Blog     string `name:"blog"`
	Company  string `name:"company"`
	Email    string `name:"email"`
	Hireable bool   `name:"hireable"`
	Location string `name:"location"`
	Name     string `name:"name"`
	internal.BaseCmd
}

func (c *UsersUpdateAuthenticatedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user")
	c.UpdateBody("bio", c.Bio)
	c.UpdateBody("blog", c.Blog)
	c.UpdateBody("company", c.Company)
	c.UpdateBody("email", c.Email)
	c.UpdateBody("hireable", c.Hireable)
	c.UpdateBody("location", c.Location)
	c.UpdateBody("name", c.Name)
	return c.DoRequest("PATCH")
}
