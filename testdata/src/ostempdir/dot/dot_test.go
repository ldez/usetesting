package dot

import (
	"errors"
	"fmt"
	. "os"
	"runtime"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	return func() {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	TempDir() // want `os\.TempDir\(\) could be replaced by <t/b>\.TempDir\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	TempDir() // want `os\.TempDir\(\) could be replaced by b\.TempDir\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	v := TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if v := TempDir(); v != "" { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						TempDir(), // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
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
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v string) {}(TempDir()) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, TempDir())) // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case TempDir() == "": // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var v string = TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = err
	_ = v
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	TempDir() // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch TempDir() { // want `os\.TempDir\(\) could be replaced by t\.TempDir\(\) in .+`
	case "":
	}
}

func foobar() {
	TempDir()
}
