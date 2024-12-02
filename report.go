package usetesting

import (
	"go/ast"
	"go/token"
	"slices"

	"golang.org/x/tools/go/analysis"
)

// because [os.CreateTemp] takes 2 args.
const nbArgCreateTemp = 2

func (a *analyzer) reportCallExpr(pass *analysis.Pass, ce *ast.CallExpr, fnInfo *FuncInfo) bool {
	if !a.osCreateTemp {
		return false
	}

	if len(ce.Args) != nbArgCreateTemp {
		return false
	}

	switch fun := ce.Fun.(type) {
	case *ast.SelectorExpr:
		if fun.Sel == nil || fun.Sel.Name != createTempName {
			return false
		}

		expr, ok := fun.X.(*ast.Ident)
		if !ok {
			return false
		}

		if expr.Name == osPkgName && isFirstArgEmptyString(ce) {
			pass.Reportf(ce.Pos(),
				`%s.%s("", ...) could be replaced by %[1]s.%[2]s(%s.%s(), ...) in %s`,
				osPkgName, createTempName, fnInfo.ArgName, tempDirName, fnInfo.Name,
			)

			return true
		}

	case *ast.Ident:
		if fun.Name != createTempName {
			return false
		}

		pkgName := getPkgNameFromType(pass, fun)

		if pkgName == osPkgName && isFirstArgEmptyString(ce) {
			pass.Reportf(ce.Pos(),
				`%s.%s("", ...) could be replaced by %[1]s.%[2]s(%s.%s(), ...) in %s`,
				osPkgName, createTempName, fnInfo.ArgName, tempDirName, fnInfo.Name,
			)

			return true
		}
	}

	return false
}

func (a *analyzer) reportSelector(pass *analysis.Pass, se *ast.SelectorExpr, fnInfo *FuncInfo) {
	if se.Sel == nil || !se.Sel.IsExported() {
		return
	}

	ident, ok := se.X.(*ast.Ident)
	if !ok {
		return
	}

	a.report(pass, se.Pos(), ident.Name, se.Sel.Name, fnInfo)
}

func (a *analyzer) reportIdent(pass *analysis.Pass, ident *ast.Ident, fnInfo *FuncInfo) {
	if !ident.IsExported() {
		return
	}

	if !slices.Contains(a.fieldNames, ident.Name) {
		return
	}

	pkgName := getPkgNameFromType(pass, ident)

	a.report(pass, ident.Pos(), pkgName, ident.Name, fnInfo)
}

//nolint:gocyclo // The complexity is expected by the number of cases to check.
func (a *analyzer) report(pass *analysis.Pass, pos token.Pos, origPkgName, origName string, fnInfo *FuncInfo) {
	switch {
	case a.osMkdirTemp && origPkgName == osPkgName && origName == mkdirTempName:
		report(pass, pos, origPkgName, origName, tempDirName, fnInfo)

	case a.osTempDir && origPkgName == osPkgName && origName == tempDirName:
		report(pass, pos, origPkgName, origName, tempDirName, fnInfo)

	case a.osSetenv && origPkgName == osPkgName && origName == setenvName:
		report(pass, pos, origPkgName, origName, setenvName, fnInfo)

	case a.geGo124 && a.osChdir && origPkgName == osPkgName && origName == chdirName:
		report(pass, pos, origPkgName, origName, chdirName, fnInfo)

	case a.geGo124 && a.contextBackground && origPkgName == contextPkgName && origName == backgroundName:
		report(pass, pos, origPkgName, origName, contextName, fnInfo)

	case a.geGo124 && a.contextTodo && origPkgName == contextPkgName && origName == todoName:
		report(pass, pos, origPkgName, origName, contextName, fnInfo)
	}
}

func report(pass *analysis.Pass, pos token.Pos, origPkgName, origName, expectName string, fnInfo *FuncInfo) {
	pass.Reportf(
		pos,
		"%s.%s() could be replaced by %s.%s() in %s",
		origPkgName, origName, fnInfo.ArgName, expectName, fnInfo.Name,
	)
}

func isFirstArgEmptyString(ce *ast.CallExpr) bool {
	bl, ok := ce.Args[0].(*ast.BasicLit)
	if !ok {
		return false
	}

	return bl.Kind == token.STRING && bl.Value == `""`
}

func getPkgNameFromType(pass *analysis.Pass, ident *ast.Ident) string {
	o := pass.TypesInfo.ObjectOf(ident)

	if o == nil || o.Pkg() == nil {
		return ""
	}

	return o.Pkg().Name()
}
