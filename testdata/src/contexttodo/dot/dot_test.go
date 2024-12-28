package dot

import (
	. "context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	return func() {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	TODO() // want `context\.TODO\(\) could be replaced by <t/b>\.Context\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	TODO() // want `context\.TODO\(\) could be replaced by b\.Context\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	ctx := TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	_ = ctx
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if ctx := TODO(); ctx != nil { // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(TODO()) // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						TODO(), // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
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
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(ctx Context) {}(TODO()) // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, TODO())) // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case TODO() == nil: // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var ctx Context = TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	_ = ctx
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, ctx any = errors.New(""), TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	_ = err
	_ = ctx
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	TODO() // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := TODO().(type) { // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch TODO() { // want `context\.TODO\(\) could be replaced by t\.Context\(\) in .+`
	case nil:
	}
}

func foobar() {
	TODO()
}
