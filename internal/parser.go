package internal

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"net/http"
	"strings"
)

func ServePkgDoc(w http.ResponseWriter, r *http.Request, pkgPath string) {
	pkgDoc, err := getPkgDoc(pkgPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("doc").Funcs(template.FuncMap{
		"fmtDecl": fmtDecl,
	}).Parse(docTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pkgDoc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getPkgDoc(pkgPath string) (*packageDoc, error) {
	// Create a new file set
	fset := token.NewFileSet()

	// Parse package file set
	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse package: %v", err)
	}

	// Assume there is only one package in the directory
	var pkg *ast.Package
	for _, p := range pkgs {
		pkg = p
		break
	}

	if pkg == nil {
		return nil, fmt.Errorf("no package found in %s", pkgPath)
	}

	docPkg := doc.New(pkg, "./", doc.AllDecls)
	return &packageDoc{
		PackageName: docPkg.Name,
		Doc:         docPkg.Doc,
		Consts:      docPkg.Consts,
		Vars:        docPkg.Vars,
		Funcs:       docPkg.Funcs,
		Types:       docPkg.Types,
		Fset:        fset,
	}, nil
}

func fmtDecl(fset *token.FileSet, decl ast.Decl) string {
	var sb strings.Builder
	printer.Fprint(&sb, fset, decl)
	return sb.String()
}
