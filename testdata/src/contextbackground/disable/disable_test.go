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
		context.Background()
	}
}

func bur(t *testing.T) func() {
	return func() {
		context.Background()
	}
}

func bir(t *testing.T) func() {
	context.Background()
	return func() {
		context.Background()
	}
}

func Test_NoName(_ *testing.T) {
	context.Background()
}

func Benchmark_ExprStmt(b *testing.B) {
	context.Background()
}

func Test_ExprStmt(t *testing.T) {
	context.Background()
}

func Test_AssignStmt(t *testing.T) {
	ctx := context.Background()
	_ = ctx
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = context.Background()
}

func Test_IfStmt(t *testing.T) {
	if ctx := context.Background(); ctx != nil {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for range 5 {
		context.Background()
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		context.Background()
	}
}

func Test_DeferStmt(t *testing.T) {
	defer context.Background()
}

func Test_CallExpr(t *testing.T) {
	t.Log(context.Background())
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						context.Background(),
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
		context.Background()
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(ctx context.Context) {}(context.Background())
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s %s", s, context.Background()))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			context.Background()
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		context.Background()
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case context.Background() == nil:
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var ctx context.Context = context.Background()
	_ = ctx
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, ctx any = errors.New(""), context.Background()
	_ = err
	_ = ctx
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				context.Background()
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		context.Background()
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					context.Background()
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		context.Background()
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	context.Background()
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := context.Background().(type) {
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch context.Background() {
	case nil:
	}
}

func foobar() {
	context.Background()
}
