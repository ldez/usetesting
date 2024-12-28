package disable

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
		os.TempDir()
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.TempDir()
	}
}

func bir(t *testing.T) func() {
	os.TempDir()
	return func() {
		os.TempDir()
	}
}

func Test_NoName(_ *testing.T) {
	os.TempDir()
}

func Benchmark_ExprStmt(b *testing.B) {
	os.TempDir()
}

func Test_ExprStmt(t *testing.T) {
	os.TempDir()
}

func Test_AssignStmt(t *testing.T) {
	v := os.TempDir()
	_ = v
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = os.TempDir()
}

func Test_IfStmt(t *testing.T) {
	if v := os.TempDir(); v != "" {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		os.TempDir()
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.TempDir()
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.TempDir()
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.TempDir())
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						os.TempDir(),
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
		os.TempDir()
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v string) {}(os.TempDir())
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, os.TempDir()))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.TempDir()
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.TempDir()
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case os.TempDir() == "":
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var v string = os.TempDir()
	_ = v
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), os.TempDir()
	_ = err
	_ = v
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.TempDir()
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.TempDir()
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.TempDir()
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		os.TempDir()
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.TempDir()
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch os.TempDir() {
	case "":
	}
}

func foobar() {
	os.TempDir()
}
