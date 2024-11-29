package nottestfiles

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func FunctionExprStmt(t *testing.T) {
	context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionAssignStmt(t *testing.T) {
	ctx := context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	_ = ctx
}

func FunctionAssignStmt_ignore_return(t *testing.T) {
	_ = context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionIfStmt(t *testing.T) {
	if ctx := context.TODO(); ctx != nil { // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}
}

func FunctionForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}
}

func FunctionDeferStmt(t *testing.T) {
	defer context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionCallExpr(t *testing.T) {
	t.Log(context.TODO()) // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionCallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						context.TODO(), // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
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
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}()
}

func FunctionGoStmt_arg(t *testing.T) {
	go func(ctx context.Context) {}(context.TODO()) // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionCallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, context.TODO())) // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionFuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+` `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
		})
	}
}

func FunctionSwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}
}

func FunctionSwitchStmt_case(t *testing.T) {
	switch {
	case context.TODO() == nil: // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
		// noop
	}
}

func FunctionDeclStmt(t *testing.T) {
	var ctx context.Context = context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	_ = ctx
}

func FunctionDeclStmt_tuple(t *testing.T) {
	var err, ctx any = errors.New(""), context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	_ = err
	_ = ctx
}

func FunctionSelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
			}
		}
	}()
}

func FunctionDeferStmt_wrap(t *testing.T) {
	defer func() {
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}()
}

func FunctionSelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
				}()
			}
		}
	}()
}

func FunctionBlockStmt(t *testing.T) {
	{
		context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	}
}

func FunctionTypeSwitchStmt(t *testing.T) {
	context.TODO() // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
}

func FunctionTypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := context.TODO().(type) { // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	case error:
		_ = v
	}
}

func FunctionSwitchStmt_Tag(t *testing.T) {
	switch context.TODO() { // want `context\.TODO\(\) could be replaced by <t/b/tb>\.Context\(\) in .+`
	case nil:
	}
}

func foobar() {
	context.TODO()
}
