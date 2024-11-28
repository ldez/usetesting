// Package usetesting It is an analyzer that detects when some calls can be replaced by methods from the testing package.
package usetesting

import (
	"go/ast"
	"slices"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	chdirName      = "Chdir"
	backgroundName = "Background"
	todoName       = "TODO"
	contextName    = "Context"
)

const (
	osPkgName      = "os"
	contextPkgName = "context"
)

// Analyzer is the usetesting analyzer.
//
//nolint:gochecknoglobals // global variables are allowed for [analysis.Analyzer].
var Analyzer = &analysis.Analyzer{
	Name:     "usetesting",
	Doc:      "reports uses of xxx instead of testing functions",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	insp, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		switch fn := node.(type) {
		case *ast.FuncDecl:
			if len(fn.Type.Params.List) < 1 {
				return
			}

			if !isTestFunction(fn.Type.Params.List[0].Type, "testing") {
				return
			}

			for _, stmt := range fn.Body.List {
				checkStmt(pass, fn.Name.Name, stmt)
			}

		case *ast.FuncLit: // TODO remove?
			if len(fn.Type.Params.List) < 1 {
				return
			}

			if !isTestFunction(fn.Type.Params.List[0].Type, "testing") {
				return
			}

			for _, stmt := range fn.Body.List {
				checkStmt(pass, "anonymous function", stmt)
			}
		}
	})

	return nil, nil
}

//nolint:funlen,gocognit,gocyclo // The complexity is expected by the number of [ast.Stmt] variants.
func checkStmt(pass *analysis.Pass, fnName string, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		checkExpr(pass, fnName, stmt.X)

	case *ast.IfStmt:
		assignStmt, ok := stmt.Init.(*ast.AssignStmt)
		if !ok {
			return
		}

		checkExpr(pass, fnName, assignStmt.Rhs[0])

	case *ast.AssignStmt:
		checkExpr(pass, fnName, stmt.Rhs[0])

	case *ast.ForStmt:
		for _, stmt := range stmt.Body.List {
			checkStmt(pass, fnName, stmt)
		}

	case *ast.DeferStmt:
		checkExpr(pass, fnName, stmt.Call)

	case *ast.RangeStmt:
		for _, stmt := range stmt.Body.List {
			checkStmt(pass, fnName, stmt)
		}

	case *ast.ReturnStmt:
		for _, result := range stmt.Results {
			checkExpr(pass, fnName, result)
		}

	case *ast.DeclStmt:
		genDecl, ok := stmt.Decl.(*ast.GenDecl)
		if !ok {
			return
		}

		valSpec, ok := genDecl.Specs[0].(*ast.ValueSpec) // TODO for?
		if !ok {
			return
		}

		for _, value := range valSpec.Values {
			checkExpr(pass, fnName, value)
		}

	case *ast.GoStmt:
		switch fun := stmt.Call.Fun.(type) {
		case *ast.FuncLit:
			for _, stmt := range fun.Body.List {
				checkStmt(pass, fnName, stmt)
			}
		default:
			checkExpr(pass, fnName, stmt.Call)
		}

	case *ast.CaseClause:
		for _, expr := range stmt.List {
			checkExpr(pass, fnName, expr)
		}

		for _, expr := range stmt.Body {
			checkStmt(pass, fnName, expr)
		}

	case *ast.SwitchStmt:
		checkExpr(pass, fnName, stmt.Tag)

		for _, s := range stmt.Body.List {
			checkStmt(pass, fnName, s)
		}

	case *ast.TypeSwitchStmt:
		if stmt.Assign != nil {
			checkStmt(pass, fnName, stmt.Assign)
		}

		for _, s := range stmt.Body.List {
			checkStmt(pass, fnName, s)
		}

	case *ast.CommClause:
		for _, s := range stmt.Body {
			checkStmt(pass, fnName, s)
		}

	case *ast.SelectStmt:
		for _, expr := range stmt.Body.List {
			checkStmt(pass, fnName, expr)
		}

	case *ast.BlockStmt:
		for _, s := range stmt.List {
			checkStmt(pass, fnName, s)
		}

	case *ast.BranchStmt, *ast.SendStmt, *ast.IncDecStmt, *ast.LabeledStmt:
		// skip

	default:
		// skip
	}
}

func checkExpr(pass *analysis.Pass, fnName string, exp ast.Expr) {
	switch expr := exp.(type) {
	case *ast.BinaryExpr:
		checkExpr(pass, fnName, expr.X)
		checkExpr(pass, fnName, expr.Y)

	case *ast.SelectorExpr:
		reportSelector(pass, expr, fnName)

	case *ast.FuncLit:
		for _, stmt := range expr.Body.List {
			checkStmt(pass, fnName, stmt)
		}

	case *ast.TypeAssertExpr:
		checkExpr(pass, fnName, expr.X)

	case *ast.CallExpr:
		for _, arg := range expr.Args {
			checkExpr(pass, fnName, arg)
		}

		checkExpr(pass, fnName, expr.Fun)

	case *ast.Ident:
		reportIdent(pass, expr, fnName)

	case *ast.BasicLit:
		// skip

	default:
		// skip
	}
}

func isTestFunction(fieldType ast.Expr, pkgName string) bool {
	switch ft := fieldType.(type) {
	case *ast.StarExpr:
		if se, ok := ft.X.(*ast.SelectorExpr); ok {
			return checkSelectorName(se, pkgName, "T", "B", "F")
		}

	case *ast.SelectorExpr:
		return checkSelectorName(ft, pkgName, "TB")
	}

	return false
}

func checkSelectorName(exp *ast.SelectorExpr, pkgName string, selectorNames ...string) bool {
	if expr, ok := exp.X.(*ast.Ident); ok {
		return pkgName == expr.Name && slices.Contains(selectorNames, exp.Sel.Name)
	}

	return false
}

func reportSelector(pass *analysis.Pass, sel *ast.SelectorExpr, fnName string) {
	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	msg := "%s.%s() could be replaced by testing.%s() in %s"

	switch {
	case pkgIdent.Name == osPkgName && sel.Sel.Name == chdirName:
		pass.Reportf(sel.Pos(), msg, pkgIdent.Name, sel.Sel.Name, chdirName, fnName)

	case pkgIdent.Name == contextPkgName && (sel.Sel.Name == todoName || sel.Sel.Name == backgroundName):
		pass.Reportf(sel.Pos(), msg, pkgIdent.Name, sel.Sel.Name, contextName, fnName)
	}
}

func reportIdent(pass *analysis.Pass, expr *ast.Ident, fnName string) {
	if expr.Name != chdirName && expr.Name != backgroundName && expr.Name != todoName {
		return
	}

	o := pass.TypesInfo.ObjectOf(expr)

	if o == nil || o.Pkg() == nil {
		return
	}

	pkgName := o.Pkg().Name()

	msg := "%s.%s() could be replaced by testing.%s() in %s"

	switch {
	case pkgName == osPkgName && expr.Name == chdirName:
		pass.Reportf(expr.Pos(), msg, pkgName, expr.Name, chdirName, fnName)

	case pkgName == contextPkgName && expr.Name != backgroundName && expr.Name != todoName:
		pass.Reportf(expr.Pos(), msg, pkgName, expr.Name, contextName, fnName)
	}
}
