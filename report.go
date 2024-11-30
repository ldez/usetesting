package usetesting

import (
	"go/ast"
	"go/token"
	"slices"

	"golang.org/x/tools/go/analysis"
)

// because [os.CreateTemp] takes 2 args.
const nbArgCreateTemp = 2

func (a *analyzer) reportCallExpr(pass *analysis.Pass, ce *ast.CallExpr, fnName string) bool {
	if !a.osCreateTemp {
		return false
	}

	if len(ce.Args) != nbArgCreateTemp {
		return false
	}

	switch fun := ce.Fun.(type) {
	case *ast.SelectorExpr:
		if fun.Sel.Name != createTempName {
			return false
		}

		expr, ok := fun.X.(*ast.Ident)
		if !ok {
			return false
		}

		if expr.Name == osPkgName {
			if isFirstArgEmptyString(ce) {
				pass.Reportf(ce.Pos(),
					`%s.%s("", ...) could be replaced by %[1]s.%[2]s(<t/b/tb>.%s(), ...) in %s`,
					osPkgName, createTempName, tempDirName, fnName,
				)

				return true
			}
		}

	case *ast.Ident:
		if fun.Name != createTempName {
			return false
		}

		pkgName := getPkgNameFromType(pass, fun)

		if pkgName == osPkgName {
			if isFirstArgEmptyString(ce) {
				pass.Reportf(ce.Pos(),
					`%s.%s("", ...) could be replaced by %[1]s.%[2]s(<t/b/tb>.%s(), ...) in %s`,
					osPkgName, createTempName, tempDirName, fnName,
				)

				return true
			}
		}
	}

	return false
}

func (a *analyzer) reportSelector(pass *analysis.Pass, sel *ast.SelectorExpr, fnName string) {
	expr, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	if !sel.Sel.IsExported() {
		return
	}

	a.report(pass, sel.Pos(), expr.Name, sel.Sel.Name, fnName)
}

func (a *analyzer) reportIdent(pass *analysis.Pass, expr *ast.Ident, fnName string) {
	if !slices.Contains(a.fieldNames, expr.Name) {
		return
	}

	if !expr.IsExported() {
		return
	}

	pkgName := getPkgNameFromType(pass, expr)

	a.report(pass, expr.Pos(), pkgName, expr.Name, fnName)
}

//nolint:gocyclo // The complexity is expected by the cases to check.
func (a *analyzer) report(pass *analysis.Pass, pos token.Pos, origPkgName, origName, fnName string) {
	switch {
	case a.osMkdirTemp && origPkgName == osPkgName && origName == mkdirTempName:
		report(pass, pos, origPkgName, origName, tempDirName, fnName)

	case a.osTempDir && origPkgName == osPkgName && origName == tempDirName:
		report(pass, pos, origPkgName, origName, tempDirName, fnName)

	case a.osSetenv && origPkgName == osPkgName && origName == setenvName:
		report(pass, pos, origPkgName, origName, setenvName, fnName)

	case a.geGo124 && a.osChdir && origPkgName == osPkgName && origName == chdirName:
		report(pass, pos, origPkgName, origName, chdirName, fnName)

	case a.geGo124 && a.contextBackground && origPkgName == contextPkgName && origName == backgroundName:
		report(pass, pos, origPkgName, origName, contextName, fnName)

	case a.geGo124 && a.contextTodo && origPkgName == contextPkgName && origName == todoName:
		report(pass, pos, origPkgName, origName, contextName, fnName)
	}
}

func report(pass *analysis.Pass, pos token.Pos, origPkgName, origName, expectName, fnName string) {
	pass.Reportf(
		pos,
		"%s.%s() could be replaced by <t/b/tb>.%s() in %s",
		origPkgName, origName, expectName, fnName,
	)
}

func isFirstArgEmptyString(ce *ast.CallExpr) bool {
	bl, ok := ce.Args[0].(*ast.BasicLit)
	if !ok {
		return false
	}

	return bl.Kind == token.STRING && bl.Value == `""`
}

func getPkgNameFromType(pass *analysis.Pass, expr *ast.Ident) string {
	o := pass.TypesInfo.ObjectOf(expr)

	if o == nil || o.Pkg() == nil {
		return ""
	}

	return o.Pkg().Name()
}
