// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type ScimCmd struct {
	ListProvisionedIdentities      ScimListProvisionedIdentitiesCmd      `cmd:"" help:"Get a list of provisioned identities - https://developer.github.com/v3/scim/#get-a-list-of-provisioned-identities"`
	GetProvisioningDetailsForUser  ScimGetProvisioningDetailsForUserCmd  `cmd:"" help:"Get provisioning details for a single user - https://developer.github.com/v3/scim/#get-provisioning-details-for-a-single-user"`
	ProvisionInviteUsers           ScimProvisionInviteUsersCmd           `cmd:"" help:"Provision and invite users - https://developer.github.com/v3/scim/#provision-and-invite-users"`
	UpdateProvisionedOrgMembership ScimUpdateProvisionedOrgMembershipCmd `cmd:"" help:"Update a provisioned organization membership - https://developer.github.com/v3/scim/#update-a-provisioned-organization-membership"`
	UpdateUserAttribute            ScimUpdateUserAttributeCmd            `cmd:"" help:"Update a user attribute - https://developer.github.com/v3/scim/#update-a-user-attribute"`
	RemoveUserFromOrg              ScimRemoveUserFromOrgCmd              `cmd:"" help:"Remove a user from the organization - https://developer.github.com/v3/scim/#remove-a-user-from-the-organization"`
}

type ScimListProvisionedIdentitiesCmd struct {
	internal.BaseCmd
	Cloud9     bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org        string `required:"" name:"org"`
	StartIndex int64  `name:"startIndex" help:"Used for pagination: the index of the first result to return."`
	Count      int64  `name:"count" help:"Used for pagination: the number of results to return."`
	Filter     string "name:\"filter\" help:\"Filters results using the equals query parameter operator (`eq`). You can filter results that are equal to `id`, `userName`, `emails`, and `external_id`. For example, to search for an identity with the `userName` Octocat, you would use this query: `?filter=userName%20eq%20\\'Octocat\\'`\""
}

func (c *ScimListProvisionedIdentitiesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("startIndex", c.StartIndex)
	c.UpdateURLQuery("count", c.Count)
	c.UpdateURLQuery("filter", c.Filter)
	return c.DoRequest("GET")
}

type ScimGetProvisioningDetailsForUserCmd struct {
	internal.BaseCmd
	Cloud9               bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org                  string `required:"" name:"org"`
	ExternalIdentityGuid string `required:"" name:"external_identity_guid"`
}

func (c *ScimGetProvisioningDetailsForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:external_identity_guid")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("external_identity_guid", c.ExternalIdentityGuid)
	return c.DoRequest("GET")
}

type ScimProvisionInviteUsersCmd struct {
	internal.BaseCmd
	Cloud9 bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org    string `required:"" name:"org"`
}

func (c *ScimProvisionInviteUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("POST")
}

type ScimUpdateProvisionedOrgMembershipCmd struct {
	internal.BaseCmd
	Cloud9               bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org                  string `required:"" name:"org"`
	ExternalIdentityGuid string `required:"" name:"external_identity_guid"`
}

func (c *ScimUpdateProvisionedOrgMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:external_identity_guid")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("external_identity_guid", c.ExternalIdentityGuid)
	return c.DoRequest("PUT")
}

type ScimUpdateUserAttributeCmd struct {
	internal.BaseCmd
	Cloud9               bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org                  string `required:"" name:"org"`
	ExternalIdentityGuid string `required:"" name:"external_identity_guid"`
}

func (c *ScimUpdateUserAttributeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:external_identity_guid")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("external_identity_guid", c.ExternalIdentityGuid)
	return c.DoRequest("PATCH")
}

type ScimRemoveUserFromOrgCmd struct {
	internal.BaseCmd
	Cloud9               bool   "name:\"cloud-9-preview\" required:\"\" help:\"**Note:** The SCIM API on GitHub is currently available for developers to preview. To access the API, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.cloud-9-preview+json+scim\n\n```\""
	Org                  string `required:"" name:"org"`
	ExternalIdentityGuid string `required:"" name:"external_identity_guid"`
}

func (c *ScimRemoveUserFromOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:external_identity_guid")
	c.UpdatePreview("cloud-9", c.Cloud9)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("external_identity_guid", c.ExternalIdentityGuid)
	return c.DoRequest("DELETE")
}