package basic

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

func Test_NoName(_ *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by <t/b>\.TempDir\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by b\.TempDir\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	v := os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if v := os.TempDir(); v != "" { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
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

func Test_GoStmt(t *testing.T) {
	go func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v string) {}(os.TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, os.TempDir())) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
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

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case os.TempDir() == "": // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var v string = os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = err
	_ = v
}

func Test_SelectStmt(t *testing.T) {
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

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
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

func Test_BlockStmt(t *testing.T) {
	{
		os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch os.TempDir() { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	case "":
	}
}

func foobar() {
	os.TempDir()
}
