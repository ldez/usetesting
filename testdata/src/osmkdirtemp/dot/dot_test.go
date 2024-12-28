package dot

import (
	"fmt"
	. "os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	return func() {
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by <t/b>\.TempDir\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by b\.TempDir\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	v, err := MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_, _ = MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if _, err := MkdirTemp("", ""); err != nil { // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		MkdirTemp("", strconv.Itoa(i)) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		MkdirTemp("", strconv.Itoa(i)) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(MkdirTemp("", "")) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf(
						MkdirTemp("", ""), // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
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
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v string, err error) {}(MkdirTemp("", "")) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf(MkdirTemp("", ""))) // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_DeclStmt(t *testing.T) {
	var v, err any = MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	_ = v
	_ = err
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	MkdirTemp("", "") // want `os\.MkdirTemp\(\) could be replaced by t\.TempDir\(\) in .+`
}

func foobar() {
	MkdirTemp("", "")
}
