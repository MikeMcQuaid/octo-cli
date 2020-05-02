// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type MarkdownCmd struct {
	Render    MarkdownRenderCmd    `cmd:""`
	RenderRaw MarkdownRenderRawCmd `cmd:""`
}

type MarkdownRenderCmd struct {
	Context string `name:"context"`
	Mode    string `name:"mode"`
	Text    string `required:"" name:"text"`
	internal.BaseCmd
}

func (c *MarkdownRenderCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/markdown")
	c.UpdateBody("context", c.Context)
	c.UpdateBody("mode", c.Mode)
	c.UpdateBody("text", c.Text)
	return c.DoRequest("POST")
}

type MarkdownRenderRawCmd struct {
	ContentType string `name:"content-type"`
	internal.BaseCmd
}

func (c *MarkdownRenderRawCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/markdown/raw")
	c.AddRequestHeader("content-type", c.ContentType)
	return c.DoRequest("POST")
}
