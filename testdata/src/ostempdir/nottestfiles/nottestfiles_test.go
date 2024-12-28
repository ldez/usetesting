package nottestfiles

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	return func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionNoName(_ *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by <t/b>\.TempDir\(\) in .+`
}

func FunctionTB(tb testing.TB) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by tb\.TempDir\(\) in .+`
}

func FunctionBench_ExprStmt(b *testing.B) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by b\.TempDir\(\) in .+`
}

func FunctionExprStmt(t *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionAssignStmt(t *testing.T) {
	v := os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func FunctionAssignStmt_ignore_return(t *testing.T) {
	_ = os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionIfStmt(t *testing.T) {
	if v := os.TempDir(); v != "" { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionDeferStmt(t *testing.T) {
	defer os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr(t *testing.T) {
	t.Log(os.TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						os.TempDir(), // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
					),
					"a",
				),
				"b",
			),
			"c",
		),
	)
}

func FunctionGoStmt(t *testing.T) {
	go func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func FunctionGoStmt_arg(t *testing.T) {
	go func(v string) {}(os.TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, os.TempDir())) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionFuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		})
	}
}

func FunctionSwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionSwitchStmt_case(t *testing.T) {
	switch {
	case os.TempDir() == "": // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// noop
	}
}

func FunctionDeclStmt(t *testing.T) {
	var v string = os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func FunctionDeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = err
	_ = v
}

func FunctionSelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
			}
		}
	}()
}

func FunctionDeferStmt_wrap(t *testing.T) {
	defer func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func FunctionSelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
				}()
			}
		}
	}()
}

func FunctionBlockStmt(t *testing.T) {
	{
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionTypeSwitchStmt(t *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionSwitchStmt_Tag(t *testing.T) {
	switch os.TempDir() { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	case "":
	}
}

func foobar() {
	os.TempDir()
}
