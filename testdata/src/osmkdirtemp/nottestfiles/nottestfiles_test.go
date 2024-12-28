package nottestfiles

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	return func() {
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionNoName(_ *testing.T) {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by <t/b>\.TempDir\(\) in .+`
}

func FunctionTB(tb testing.TB) {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by tb\.TempDir\(\) in .+`
}

func FunctionBench_ExprStmt(b *testing.B) {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by b\.TempDir\(\) in .+`
}

func FunctionExprStmt(t *testing.T) {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionAssignStmt(t *testing.T) {
	v, err := os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
	_ = err
}

func FunctionAssignStmt_ignore_return(t *testing.T) {
	_, _ = os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionIfStmt(t *testing.T) {
	if _, err := os.MkdirTemp("", ""); err != nil { // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.MkdirTemp("", strconv.Itoa(i)) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.MkdirTemp("", strconv.Itoa(i)) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionDeferStmt(t *testing.T) {
	defer os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr(t *testing.T) {
	t.Log(os.MkdirTemp("", "")) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf(
						os.MkdirTemp("", ""), // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
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
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func FunctionGoStmt_arg(t *testing.T) {
	go func(v string, err error) {}(os.MkdirTemp("", "")) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionCallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf(os.MkdirTemp("", ""))) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func FunctionFuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
		})
	}
}

func FunctionSwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionDeclStmt(t *testing.T) {
	var v, err any = os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
	_ = err
}

func FunctionSelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
			}
		}
	}()
}

func FunctionDeferStmt_wrap(t *testing.T) {
	defer func() {
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func FunctionSelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
				}()
			}
		}
	}()
}

func FunctionBlockStmt(t *testing.T) {
	{
		os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func FunctionTypeSwitchStmt(t *testing.T) {
	os.MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func foobar() {
	os.MkdirTemp("", "")
}
