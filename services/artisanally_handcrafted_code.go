package services

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-github-cli/go-github-cli/internal"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

type baseCmd struct {
	isValueSetMap map[string]bool
	url           url.URL
	reqBody       *map[string]interface{}
	Token         string `env:"GITHUB_TOKEN" required:""`
	APIBaseURL    string `env:"GITHUB_API_BASE_URL" default:"https://api.github.com"`
	RawOutput     bool   `help:"don't format json output."`
	Format        string `help:"format output with a go template"`
}

func (c *baseCmd) isValueSet(valueName string) bool {
	return c.isValueSetMap[valueName]
}

func (c *baseCmd) updateBody(flagName string, value interface{}) {
	if c.reqBody == nil {
		c.reqBody = &map[string]interface{}{}
	}
	if c.isValueSet(flagName) {
		b := *c.reqBody
		b[flagName] = value
		c.reqBody = &b
	}
}

func (c *baseCmd) updateURLPath(valName string, value interface{}) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	case int64:
		strVal = strconv.FormatInt(v, 10)
	case nil:
		strVal = ""
	}
	c.url.Path = strings.Replace(c.url.Path, ":"+valName, strVal, 1)
}

func (c *baseCmd) updateURLQuery(paramName string, value interface{}) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	case int64:
		strVal = strconv.FormatInt(v, 10)
	}
	if c.isValueSet(paramName) {
		query := c.url.Query()
		query.Add(paramName, strVal)
		c.url.RawQuery = query.Encode()
	}
}

func (c *baseCmd) newRequest(method string) (*http.Request, error) {
	u := strings.Join([]string{
		strings.TrimSuffix(c.APIBaseURL, "/"),
		strings.TrimPrefix(c.url.String(), "/"),
	}, "/")
	var buf io.ReadWriter
	if c.reqBody != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(c.reqBody)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}
	if c.reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	return req, nil
}

var transportWrapper interface {
	SetTransport(t http.RoundTripper)
	http.RoundTripper
}

func (c *baseCmd) httpClient() *http.Client {
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token}))
	if transportWrapper != nil {
		transportWrapper.SetTransport(tc.Transport)
		tc.Transport = transportWrapper
	}
	return tc
}

func (c *baseCmd) doRequest(method string) error {
	req, err := c.newRequest(method)
	if err != nil {
		return err
	}

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}

	return internal.OutputResult(resp, c.RawOutput, c.Format, stdout)
}