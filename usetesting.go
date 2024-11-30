// Package usetesting It is an analyzer that detects when some calls can be replaced by methods from the testing package.
package usetesting

import (
	"go/ast"
	"go/build"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	chdirName      = "Chdir"
	mkdirTempName  = "MkdirTemp"
	createTempName = "CreateTemp"
	setenvName     = "Setenv"
	tempDirName    = "TempDir"
	backgroundName = "Background"
	todoName       = "TODO"
	contextName    = "Context"
)

const (
	osPkgName      = "os"
	contextPkgName = "context"
	testingPkgName = "testing"
)

// analyzer is the UseTesting linter.
type analyzer struct {
	contextBackground bool
	contextTodo       bool
	osChdir           bool
	osMkdirTemp       bool
	osTempDir         bool
	osSetenv          bool
	osCreateTemp      bool

	fieldNames []string

	skipGoVersionDetection bool
	geGo124                bool
}

// NewAnalyzer create a new UseTesting.
func NewAnalyzer() *analysis.Analyzer {
	_, skip := os.LookupEnv("USETESTING_SKIP_GO_VERSION_CHECK") // TODO should be removed when go1.25 will be released.

	l := &analyzer{
		fieldNames: []string{
			chdirName,
			mkdirTempName,
			tempDirName,
			setenvName,
			backgroundName,
			todoName,
			createTempName,
		},
		skipGoVersionDetection: skip,
	}

	a := &analysis.Analyzer{
		Name:     "usetesting",
		Doc:      "Reports uses of functions with replacement inside the testing package.",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      l.run,
	}

	a.Flags.BoolVar(&l.contextBackground, "contextbackground", true, "Enable/disable context.Background() detections")
	a.Flags.BoolVar(&l.contextTodo, "contexttodo", true, "Enable/disable context.TODO() detections")
	a.Flags.BoolVar(&l.osChdir, "oschdir", true, "Enable/disable os.Chdir() detections")
	a.Flags.BoolVar(&l.osMkdirTemp, "osmkdirtemp", true, "Enable/disable os.MkdirTemp() detections")
	a.Flags.BoolVar(&l.osSetenv, "ossetenv", false, "Enable/disable os.Setenv() detections")
	a.Flags.BoolVar(&l.osTempDir, "ostempdir", false, "Enable/disable os.TempDir() detections")
	a.Flags.BoolVar(&l.osCreateTemp, "oscreatetemp", true, `Enable/disable os.CreateTemp("", ...) detections`)

	return a
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	if !a.osChdir && !a.contextBackground && !a.contextTodo && !a.osMkdirTemp && !a.osSetenv {
		return nil, nil
	}

	a.geGo124 = a.isGoSupported(pass)

	insp, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		switch fn := node.(type) {
		case *ast.FuncDecl:
			a.checkFunc(pass, fn.Type, fn.Body, fn.Name.Name)

		case *ast.FuncLit:
			a.checkFunc(pass, fn.Type, fn.Body, "anonymous function")
		}
	})

	return nil, nil
}

func (a *analyzer) isGoSupported(pass *analysis.Pass) bool {
	if a.skipGoVersionDetection {
		return true
	}

	// Prior to go1.22, versions.FileVersion returns only the toolchain version,
	// which is of no use to us,
	// so disable this analyzer on earlier versions.
	if !slices.Contains(build.Default.ReleaseTags, "go1.22") {
		return false
	}

	pkgVersion := pass.Pkg.GoVersion()
	if pkgVersion == "" {
		// Empty means Go devel.
		return true
	}

	vParts := strings.Split(strings.TrimPrefix(pkgVersion, "go"), ".")

	v, err := strconv.Atoi(strings.Join(vParts[:2], ""))
	if err != nil {
		v = 116
	}

	return v >= 124
}

func (a *analyzer) checkFunc(pass *analysis.Pass, ft *ast.FuncType, block *ast.BlockStmt, fnName string) {
	if len(ft.Params.List) < 1 {
		return
	}

	if !isTestFunction(ft.Params.List[0].Type, testingPkgName) {
		return
	}

	checkStmts(a, pass, fnName, block.List)
}

//nolint:funlen // The complexity is expected by the number of [ast.Stmt] variants.
func (a *analyzer) checkStmt(pass *analysis.Pass, fnName string, stmt ast.Stmt) {
	if stmt == nil {
		return
	}

	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		a.checkExpr(pass, fnName, stmt.X)

	case *ast.IfStmt:
		a.checkStmt(pass, fnName, stmt.Init)

	case *ast.AssignStmt:
		a.checkExpr(pass, fnName, stmt.Rhs[0])

	case *ast.ForStmt:
		a.checkStmt(pass, fnName, stmt.Body)

	case *ast.DeferStmt:
		a.checkExpr(pass, fnName, stmt.Call)

	case *ast.RangeStmt:
		a.checkStmt(pass, fnName, stmt.Body)

	case *ast.ReturnStmt:
		checkExprs(a, pass, fnName, stmt.Results)

	case *ast.DeclStmt:
		genDecl, ok := stmt.Decl.(*ast.GenDecl)
		if !ok {
			return
		}

		valSpec, ok := genDecl.Specs[0].(*ast.ValueSpec) // TODO for?
		if !ok {
			return
		}

		checkExprs(a, pass, fnName, valSpec.Values)

	case *ast.GoStmt:
		a.checkExpr(pass, fnName, stmt.Call)

	case *ast.CaseClause:
		checkExprs(a, pass, fnName, stmt.List)
		checkStmts(a, pass, fnName, stmt.Body)

	case *ast.SwitchStmt:
		a.checkExpr(pass, fnName, stmt.Tag)
		a.checkStmt(pass, fnName, stmt.Body)

	case *ast.TypeSwitchStmt:
		a.checkStmt(pass, fnName, stmt.Assign)
		a.checkStmt(pass, fnName, stmt.Body)

	case *ast.CommClause:
		checkStmts(a, pass, fnName, stmt.Body)

	case *ast.SelectStmt:
		a.checkStmt(pass, fnName, stmt.Body)

	case *ast.BlockStmt:
		checkStmts(a, pass, fnName, stmt.List)

	case *ast.BranchStmt, *ast.SendStmt, *ast.IncDecStmt, *ast.LabeledStmt:
		// skip

	default:
		// skip
	}
}

func (a *analyzer) checkExpr(pass *analysis.Pass, fnName string, exp ast.Expr) {
	switch expr := exp.(type) {
	case *ast.BinaryExpr:
		a.checkExpr(pass, fnName, expr.X)
		a.checkExpr(pass, fnName, expr.Y)

	case *ast.SelectorExpr:
		a.reportSelector(pass, expr, fnName)

	case *ast.FuncLit:
		for _, stmt := range expr.Body.List {
			a.checkStmt(pass, fnName, stmt)
		}

	case *ast.TypeAssertExpr:
		a.checkExpr(pass, fnName, expr.X)

	case *ast.CallExpr:
		if a.reportCallExpr(pass, expr, fnName) {
			return
		}

		for _, arg := range expr.Args {
			a.checkExpr(pass, fnName, arg)
		}

		a.checkExpr(pass, fnName, expr.Fun)

	case *ast.Ident:
		a.reportIdent(pass, expr, fnName)

	case *ast.BasicLit:
		// skip

	default:
		// skip
	}
}

func checkStmts[T ast.Stmt](a *analyzer, pass *analysis.Pass, fnName string, stmts []T) {
	for _, stmt := range stmts {
		a.checkStmt(pass, fnName, stmt)
	}
}

func checkExprs(a *analyzer, pass *analysis.Pass, fnName string, exprs []ast.Expr) {
	for _, expr := range exprs {
		a.checkExpr(pass, fnName, expr)
	}
}

func isTestFunction(argType ast.Expr, pkgName string) bool {
	switch ft := argType.(type) {
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
