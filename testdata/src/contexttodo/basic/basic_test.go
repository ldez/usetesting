package basic

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func Test_ExprStmt(t *testing.T) {
	context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	ctx := context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	_ = ctx
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if ctx := context.TODO(); ctx != nil { // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(context.TODO()) // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						context.TODO(), // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
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
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(ctx context.Context) {}(context.TODO()) // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, context.TODO())) // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+` `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case context.TODO() == nil: // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var ctx context.Context = context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	_ = ctx
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, ctx any = errors.New(""), context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	_ = err
	_ = ctx
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	context.TODO() // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := context.TODO().(type) { // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch context.TODO() { // want `context\.TODO\(\) could be replaced by testing\.Context\(\) in .+`
	case nil:
	}
}

func foobar() {
	context.TODO()
}