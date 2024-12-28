package disable

import (
	"os"
	"runtime"
	"strconv"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		os.CreateTemp("", "")
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.CreateTemp("", "")
	}
}

func bir(t *testing.T) func() {
	os.CreateTemp("", "")
	return func() {
		os.CreateTemp("", "")
	}
}

func Test_NoName(_ *testing.T) {
	os.CreateTemp("", "")
}

func Benchmark_ExprStmt(b *testing.B) {
	os.CreateTemp("", "")
}

func Test_ExprStmt(t *testing.T) {
	os.CreateTemp("", "")
	os.CreateTemp("", "xx")
	os.CreateTemp(os.TempDir(), "xx")
	os.CreateTemp(t.TempDir(), "xx")
}

func Test_AssignStmt(t *testing.T) {
	f, err := os.CreateTemp("", "")
	_ = err
	_ = f
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_, _ = os.CreateTemp("", "")
}

func Test_IfStmt(t *testing.T) {
	if _, err := os.CreateTemp("", ""); err != nil {
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.CreateTemp("", strconv.Itoa(i))
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.CreateTemp("", strconv.Itoa(i))
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.CreateTemp("", "")
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.CreateTemp("", ""))
}

func Test_GoStmt(t *testing.T) {
	go func() {
		os.CreateTemp("", "")
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v *os.File, err error) {}(os.CreateTemp("", ""))
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.CreateTemp("", "")
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.CreateTemp("", "")
	}
}

func Test_DeclStmt(t *testing.T) {
	var f, err any = os.CreateTemp("", "")
	_ = err
	_ = f
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.CreateTemp("", "")
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.CreateTemp("", "")
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.CreateTemp("", "")
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		os.CreateTemp("", "")
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.CreateTemp("", "")
}

func foobar() {
	os.CreateTemp("", "")
}
