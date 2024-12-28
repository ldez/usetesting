package dot

import (
	"errors"
	"fmt"
	. "os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	return func() {
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	Chdir("") // want `os\.Chdir\(\) could be replaced by <t/b>\.Chdir\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	Chdir("") // want `os\.Chdir\(\) could be replaced by b\.Chdir\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	err := Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if err := Chdir(""); err != nil { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		Chdir(strconv.Itoa(i)) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		Chdir(strconv.Itoa(i)) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(Chdir("")) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						Chdir(""), // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
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
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(err error) {}(Chdir("")) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s", Chdir(s))) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case Chdir("") == nil: // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var err error = Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, r error = errors.New(""), Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
	_ = r
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := Chdir("").(type) { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch Chdir("") { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	case errors.New(""):
	}
}

func foobar() {
	Chdir("")
}
