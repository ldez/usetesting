package disable

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
		os.MkdirTemp("", "")
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.MkdirTemp("", "")
	}
}

func bir(t *testing.T) func() {
	os.MkdirTemp("", "")
	return func() {
		os.MkdirTemp("", "")
	}
}

func Test_NoName(_ *testing.T) {
	os.MkdirTemp("", "")
}

func Benchmark_ExprStmt(b *testing.B) {
	os.MkdirTemp("", "")
}

func Test_ExprStmt(t *testing.T) {
	os.MkdirTemp("", "")
}

func Test_AssignStmt(t *testing.T) {
	v, err := os.MkdirTemp("", "")
	_ = v
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_, _ = os.MkdirTemp("", "")
}

func Test_IfStmt(t *testing.T) {
	if _, err := os.MkdirTemp("", ""); err != nil {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.MkdirTemp("", strconv.Itoa(i))
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.MkdirTemp("", strconv.Itoa(i))
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.MkdirTemp("", "")
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.MkdirTemp("", ""))
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf(
						os.MkdirTemp("", ""),
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
		os.MkdirTemp("", "")
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v string, err error) {}(os.MkdirTemp("", ""))
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf(os.MkdirTemp("", "")))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.MkdirTemp("", "")
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.MkdirTemp("", "")
	}
}

func Test_DeclStmt(t *testing.T) {
	var v, err any = os.MkdirTemp("", "")
	_ = v
	_ = err
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.MkdirTemp("", "")
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.MkdirTemp("", "")
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.MkdirTemp("", "")
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		os.MkdirTemp("", "")
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.MkdirTemp("", "")
}

func foobar() {
	os.MkdirTemp("", "")
}
