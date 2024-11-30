package nottestfiles

import (
	"os"
	"runtime"
	"strconv"
	"testing"
)

func FunctionNoName(_ *testing.T) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(<t/b/f>\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionTB(tb testing.TB) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(tb\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionBench_ExprStmt(b *testing.B) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(b\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionExprStmt(t *testing.T) {
	os.CreateTemp("", "")   // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	os.CreateTemp("", "xx") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	os.CreateTemp(os.TempDir(), "xx")
	os.CreateTemp(t.TempDir(), "xx")
}

func FunctionAssignStmt(t *testing.T) {
	f, err := os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func FunctionAssignStmt_ignore_return(t *testing.T) {
	_, _ = os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionIfStmt(t *testing.T) {
	if _, err := os.CreateTemp("", ""); err != nil { // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func FunctionForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func FunctionDeferStmt(t *testing.T) {
	defer os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionCallExpr(t *testing.T) {
	t.Log(os.CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionGoStmt(t *testing.T) {
	go func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func FunctionGoStmt_arg(t *testing.T) {
	go func(v *os.File, err error) {}(os.CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func FunctionFuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+` `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		})
	}
}

func FunctionSwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func FunctionDeclStmt(t *testing.T) {
	var f, err any = os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func FunctionSelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
			}
		}
	}()
}

func FunctionDeferStmt_wrap(t *testing.T) {
	defer func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func FunctionSelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
				}()
			}
		}
	}()
}

func FunctionBlockStmt(t *testing.T) {
	{
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func FunctionTypeSwitchStmt(t *testing.T) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func foobar() {
	os.CreateTemp("", "")
}
