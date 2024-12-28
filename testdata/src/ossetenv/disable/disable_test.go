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
		os.Setenv("foo", "bar")
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.Setenv("foo", "bar")
	}
}

func bir(t *testing.T) func() {
	os.Setenv("foo", "bar")
	return func() {
		os.Setenv("foo", "bar")
	}
}

func Test_NoName(_ *testing.T) {
	os.Setenv("", "")
}

func Benchmark_ExprStmt(b *testing.B) {
	os.Setenv("", "")
}

func Test_ExprStmt(t *testing.T) {
	os.Setenv("", "")
}

func Test_AssignStmt(t *testing.T) {
	err := os.Setenv("", "")
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = os.Setenv("", "")
}

func Test_IfStmt(t *testing.T) {
	if err := os.Setenv("", ""); err != nil {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		os.Setenv("", "")
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.Setenv("", "")
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.Setenv("", "")
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.Setenv("", ""))
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						os.Setenv("", ""),
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
		os.Setenv("", "")
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(err error) {}(os.Setenv("", ""))
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, os.Setenv("", "")))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.Setenv("", "")
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.Setenv("", "")
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case os.Setenv("", "") == nil:
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var err error = os.Setenv("", "")
	_ = err
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), os.Setenv("", "")
	_ = err
	_ = v
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.Setenv("", "")
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.Setenv("", "")
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.Setenv("", "")
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		os.Setenv("", "")
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.Setenv("", "")
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := os.Setenv("", "").(type) {
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch os.Setenv("", "") {
	case nil:
	}
}

func foobar() {
	os.Setenv("", "")
}
