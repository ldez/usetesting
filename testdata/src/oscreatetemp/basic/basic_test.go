package basic

import (
	"os"
	"runtime"
	"strconv"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func bir(t *testing.T) func() {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	return func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(<t/b>\.TempDir\(\), \.\.\.\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(b\.TempDir\(\), \.\.\.\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	os.CreateTemp("", "")   // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	os.CreateTemp("", "xx") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	os.CreateTemp(os.TempDir(), "xx")
	os.CreateTemp(t.TempDir(), "xx")
}

func Test_AssignStmt(t *testing.T) {
	f, err := os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_, _ = os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if _, err := os.CreateTemp("", ""); err != nil { // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_GoStmt(t *testing.T) {
	go func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v *os.File, err error) {}(os.CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_DeclStmt(t *testing.T) {
	var f, err any = os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func Test_SelectStmt(t *testing.T) {
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

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
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

func Test_BlockStmt(t *testing.T) {
	{
		os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func foobar() {
	os.CreateTemp("", "")
}
