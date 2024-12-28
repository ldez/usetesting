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
		Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	return func() {
		Setenv("foo", "bar") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	Setenv("", "") // want `os\.Setenv\(\) could be replaced by <t/b>\.Setenv\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	Setenv("", "") // want `os\.Setenv\(\) could be replaced by b\.Setenv\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	err := Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if err := Setenv("", ""); err != nil { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(Setenv("", "")) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						Setenv("", ""), // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
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
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(err error) {}(Setenv("", "")) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, Setenv("", ""))) // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case Setenv("", "") == nil: // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var err error = Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, v any = errors.New(""), Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	_ = err
	_ = v
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	Setenv("", "") // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := Setenv("", "").(type) { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch Setenv("", "") { // want `os\.Setenv\(\) could be replaced by t\.Setenv\(\) in .+`
	case nil:
	}
}

func foobar() {
	Setenv("", "")
}
