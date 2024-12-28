package disable

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		context.TODO()
	}
}

func bur(t *testing.T) func() {
	return func() {
		context.TODO()
	}
}

func bir(t *testing.T) func() {
	context.TODO()
	return func() {
		context.TODO()
	}
}

func Test_NoName(_ *testing.T) {
	context.TODO()
}

func Benchmark_ExprStmt(b *testing.B) {
	context.TODO()
}

func Test_ExprStmt(t *testing.T) {
	context.TODO()
}

func Test_AssignStmt(t *testing.T) {
	ctx := context.TODO()
	_ = ctx
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = context.TODO()
}

func Test_IfStmt(t *testing.T) {
	if ctx := context.TODO(); ctx != nil {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		context.TODO()
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		context.TODO()
	}
}

func Test_DeferStmt(t *testing.T) {
	defer context.TODO()
}

func Test_CallExpr(t *testing.T) {
	t.Log(context.TODO())
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						context.TODO(),
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
		context.TODO()
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(ctx context.Context) {}(context.TODO())
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, context.TODO()))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			context.TODO()
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		context.TODO()
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case context.TODO() == nil:
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var ctx context.Context = context.TODO()
	_ = ctx
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, ctx any = errors.New(""), context.TODO()
	_ = err
	_ = ctx
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				context.TODO()
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		context.TODO()
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					context.TODO()
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		context.TODO()
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	context.TODO()
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := context.TODO().(type) {
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch context.TODO() {
	case nil:
	}
}

func foobar() {
	context.TODO()
}
