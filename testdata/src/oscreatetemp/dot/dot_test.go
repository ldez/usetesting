package dot

import (
	. "os"
	"runtime"
	"strconv"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func bir(t *testing.T) func() {
	CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	return func() {
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(<t/b>\.TempDir\(\), \.\.\.\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(b\.TempDir\(\), \.\.\.\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	CreateTemp("", "")   // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	CreateTemp("", "xx") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	CreateTemp(TempDir(), "xx")
	CreateTemp(t.TempDir(), "xx")
}

func Test_AssignStmt(t *testing.T) {
	f, err := CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_, _ = CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if _, err := CreateTemp("", ""); err != nil { // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateTemp("", strconv.Itoa(i)) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_GoStmt(t *testing.T) {
	go func() {
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(v *File, err error) {}(CreateTemp("", "")) // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_DeclStmt(t *testing.T) {
	var f, err any = CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	_ = err
	_ = f
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	CreateTemp("", "") // want `os\.CreateTemp\("", \.\.\.\) could be replaced by os\.CreateTemp\(t\.TempDir\(\), \.\.\.\) in .+`
}

func foobar() {
	CreateTemp("", "")
}
