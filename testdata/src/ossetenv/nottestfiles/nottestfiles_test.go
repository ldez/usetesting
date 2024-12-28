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
		os.Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	os.Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	return func() {
		os.Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func FunctionNoName(_ *testing.T) {
	os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by <t/b>\.Setenv\(\) in .+`
}

func FunctionTB(tb testing.TB) {
	os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by tb\.Setenv\(\) in .+`
}

func FunctionBench_ExprStmt(b *testing.B) {
	os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by b\.Setenv\(\) in .+`
}

func FunctionExprStmt(t *testing.T) {
	os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionAssignStmt(t *testing.T) {
	err := os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
}

func FunctionAssignStmt_ignore_return(t *testing.T) {
	_ = os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionIfStmt(t *testing.T) {
	if err := os.Setenv("", ""); err != nil { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func FunctionForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func FunctionDeferStmt(t *testing.T) {
	defer os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionCallExpr(t *testing.T) {
	t.Log(os.Setenv("", "")) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionCallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						os.Setenv("", ""), // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
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
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}()
}

func FunctionGoStmt_arg(t *testing.T) {
	go func(err error) {}(os.Setenv("", "")) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionCallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, os.Setenv("", ""))) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionFuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		})
	}
}

func FunctionSwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func FunctionSwitchStmt_case(t *testing.T) {
	switch {
	case os.Setenv("", "") == nil: // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		// noop
	}
}

func FunctionDeclStmt(t *testing.T) {
	var err error = os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
}

func FunctionDeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
	_ = v
}

func FunctionSelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
			}
		}
	}()
}

func FunctionDeferStmt_wrap(t *testing.T) {
	defer func() {
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}()
}

func FunctionSelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
				}()
			}
		}
	}()
}

func FunctionBlockStmt(t *testing.T) {
	{
		os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func FunctionTypeSwitchStmt(t *testing.T) {
	os.Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func FunctionTypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := os.Setenv("", "").(type) { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	case error:
		_ = v
	}
}

func FunctionSwitchStmt_Tag(t *testing.T) {
	switch os.Setenv("", "") { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	case nil:
	}
}

func foobar() {
	os.Setenv("", "")
}
