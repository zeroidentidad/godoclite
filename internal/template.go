package internal

import (
	"go/doc"
	"go/token"
)

// packageDoc hold package documentation data
type packageDoc struct {
	PackageName string
	Doc         string
	Consts      []*doc.Value
	Vars        []*doc.Value
	Funcs       []*doc.Func
	Types       []*doc.Type
	Fset        *token.FileSet
}

const docTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Doc: {{.PackageName}}</title>
	<style>
		body { font-family: Arial, sans-serif; line-height: 1.6; }
		h1, h2, h3 { color: #333; }
		pre { background: #f4f4f4; padding: 10px; border: 1px solid #ddd; }
	</style>
</head>
<body>
	<h1>Package {{.PackageName}}</h1>
	<p>{{.Doc}}</p>

  {{if .Consts}}
    <h2>Constants</h2>
    {{range .Consts}}
    <pre>{{fmtDecl $.Fset .Decl}}</pre>
    <p>{{.Doc}}</p>
    {{end}}
  {{end}}

  {{if .Vars}}
    <h2>Variables</h2>
    {{range .Vars}}
    <pre>{{fmtDecl $.Fset .Decl}}</pre>
    <p>{{.Doc}}</p>
    {{end}}
  {{end}}

  {{if .Funcs}}
    <h2>Functions</h2>
    {{range .Funcs}}
    <h3>{{.Name}}</h3>
    <pre>{{fmtDecl $.Fset .Decl}}</pre>
    <p>{{.Doc}}</p>
    {{end}}
  {{end}}

  {{if .Types}}
    <h2>Types</h2>
    {{range .Types}}
    <h3>{{.Name}}</h3>
    <pre>{{fmtDecl $.Fset .Decl}}</pre>
    <p>{{.Doc}}</p>
      {{if .Methods}}
      <h4>Methods</h4>
        {{range .Methods}}
        <h5>{{.Name}}</h5>
        <pre>{{fmtDecl $.Fset .Decl}}</pre>
        <p>{{.Doc}}</p>
        {{end}}
      {{end}}
    {{end}}
  {{end}}
</body>
</html>
`
