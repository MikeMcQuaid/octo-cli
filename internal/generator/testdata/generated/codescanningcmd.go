// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type CodeScanningCmd struct {
	GetAlert          CodeScanningGetAlertCmd          `cmd:""`
	ListAlertsForRepo CodeScanningListAlertsForRepoCmd `cmd:""`
}

type CodeScanningGetAlertCmd struct {
	AlertId int64  `required:"" name:"alert_id"`
	Owner   string `name:"owner"`
	Repo    string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *CodeScanningGetAlertCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/code-scanning/alerts/{alert_id}")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("alert_id", c.AlertId)
	return c.DoRequest("GET")
}

type CodeScanningListAlertsForRepoCmd struct {
	Owner string `name:"owner"`
	Ref   string `name:"ref"`
	Repo  string `required:"" name:"repo"`
	State string `name:"state"`
	internal.BaseCmd
}

func (c *CodeScanningListAlertsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/code-scanning/alerts")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("ref", c.Ref)
	return c.DoRequest("GET")
}
