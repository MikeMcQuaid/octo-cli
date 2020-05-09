package codegen

import (
	"bytes"
	"go/format"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/spf13/afero"
)

var tmpl = template.Must(template.New("").Parse(tmplFileTmpl))

func init() {
	template.Must(tmpl.New("RunMethodParam").Parse(tmplRunMethodParam))
	template.Must(tmpl.New("RunMethod").Parse(tmplRunMethod))
	template.Must(tmpl.New("StructType").Parse(tmplStructType))
	template.Must(tmpl.New("SvcTmpl").Parse(tmplSvcTmpl))
	template.Must(tmpl.New("CmdHelps").Parse(tmplCmdHelps))
	template.Must(tmpl.New("FlagHelps").Parse(tmplFlagHelps))
}

type RunMethodParam struct {
	Name         string
	ValueField   string
	UpdateMethod string
}

// language=GoTemplate
const tmplRunMethodParam = `
c.{{.UpdateMethod}}("{{.Name}}", c.{{.ValueField}})`

type RunMethod struct {
	ReceiverName string
	Method       string
	URLPath      string
	Params       []RunMethodParam
}

// language=GoTemplate
const tmplRunMethod = `
func (c *{{.ReceiverName}}) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("{{.URLPath}}"){{range .Params}}{{template "RunMethodParam" .}}{{end}}
	return c.DoRequest("{{.Method}}")
}

`

// StructField is one field in a StructTmplHelper
type StructField struct {
	Name string
	Type string
	Tags *structtag.Tags
}

type StructTmplHelper struct {
	Name   string
	Fields []StructField
}

// language=GoTemplate
const tmplStructType = `type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} {{if .Tags}}{{printf "%#q" .Tags}} {{end}}{{end}}
}`

type CmdStructAndMethod struct {
	CmdStruct StructTmplHelper
	RunMethod RunMethod
}

type SvcTmpl struct {
	SvcStruct           StructTmplHelper
	CmdStructAndMethods []CmdStructAndMethod
}

// language=GoTemplate
const tmplSvcTmpl = `{{if .SvcStruct}}{{template "StructType" .SvcStruct}}{{end}}
	{{range .CmdStructAndMethods}}
	{{if .CmdStruct}}{{template "StructType" .CmdStruct}}{{end}}
	{{if .RunMethod}}{{template "RunMethod" .RunMethod}}{{end}}
	{{end}}
`

type FileTmpl struct {
	CmdHelps       map[string]map[string]string
	FlagHelps      map[string]map[string]map[string]string
	PrimaryStructs []StructTmplHelper
	SvcTmpls       []SvcTmpl
}

// language=GoTemplate
const tmplCmdHelps = `
{{if .}}
  var CmdHelps = map[string]map[string]string{
  {{range $topCmd, $topCmdVals := .}}"{{$topCmd}}": {
  {{range $cmd, $help := $topCmdVals}}"{{$cmd}}": {{printf "%q" $help}},
  {{end}}
  },
  {{end}}
  }
{{end}}
`

// language=GoTemplate
const tmplFlagHelps = `
{{if .}}
  var FlagHelps = map[string]map[string]map[string]string{
  {{range $topCmd, $topCmdVals := .}}"{{$topCmd}}": {
  {{range $cmd, $flagHelps := $topCmdVals}}"{{$cmd}}": {
  {{range $flag, $help := $flagHelps}}"{{$flag}}": {{printf "%q" $help}},
  {{end}}
  },
  {{end}}
  },
  {{end}}
  }
{{end}}
`

// language=GoTemplate
const tmplFileTmpl = `// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

{{if .SvcTmpls}}import "github.com/octo-cli/octo-cli/internal"{{end}}

{{range .PrimaryStructs}}
    {{template "StructType" .}}
{{end}}

{{range .SvcTmpls}}{{template "SvcTmpl" .}}{{end}}

{{template "CmdHelps" .CmdHelps}}

{{template "FlagHelps" .FlagHelps}}

`

func WriteFiles(files map[string]FileTmpl, outputPath string, fs afero.Fs) error {
	for name, fileTmpl := range files {
		err := writeGoFile(name, fileTmpl, outputPath, fs)
		if err != nil {
			return err
		}
	}
	return nil
}

//writeGoFile executes the named template and does the equivalent of `go fmt` and `goimports` on the output
func writeGoFile(filename string, fileTmpl FileTmpl, path string, fs afero.Fs) error {
	out, err := generateGoFile(fileTmpl)
	if err != nil {
		return err
	}
	fl := filepath.Join(path, filename)
	return afero.WriteFile(fs, fl, out, 0644)
}

func generateGoFile(fileTmpl FileTmpl) ([]byte, error) {
	for _, svcTmpl := range fileTmpl.SvcTmpls {
		tmplSorting(&svcTmpl)
	}
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "", fileTmpl)
	if err != nil {
		return nil, err
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return out, nil
}

func sortCmdStructFields(fields []StructField) {
	if len(fields) == 0 {
		return
	}
	newFields := make([]StructField, 0, len(fields))
	holdouts := make([]StructField, 0, len(fields))
	for _, field := range fields {
		if field.Name == "" {
			holdouts = append(holdouts, field)
			continue
		}
		newFields = append(newFields, field)
	}
	sort.Slice(newFields, func(i, j int) bool {
		return newFields[i].Name < newFields[j].Name
	})
	sort.Slice(holdouts, func(i, j int) bool {
		return holdouts[i].Type < holdouts[j].Type
	})
	newFields = append(newFields, holdouts...)
	copy(fields, newFields)
}

func tmplSorting(svcTmpl *SvcTmpl) {
	sort.Slice(svcTmpl.SvcStruct.Fields, func(i, j int) bool {
		return svcTmpl.SvcStruct.Fields[i].Name < svcTmpl.SvcStruct.Fields[j].Name
	})
	for _, csm := range svcTmpl.CmdStructAndMethods {
		sortCmdStructFields(csm.CmdStruct.Fields)

		sort.Slice(csm.RunMethod.Params, func(i, j int) bool {
			return csm.RunMethod.Params[i].Name < csm.RunMethod.Params[j].Name
		})
	}
	sort.Slice(svcTmpl.CmdStructAndMethods, func(i, j int) bool {
		return svcTmpl.CmdStructAndMethods[i].CmdStruct.Name < svcTmpl.CmdStructAndMethods[j].CmdStruct.Name
	})
}
