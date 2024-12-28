package basic

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func bar() func(t *testing.T) {
	return func(t *testing.T) {
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func bur(t *testing.T) func() {
	return func() {
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func bir(t *testing.T) func() {
	os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	return func() {
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_NoName(_ *testing.T) {
	os.Chdir("") // want `os\.Chdir\(\) could be replaced by <t/b>\.Chdir\(\) in .+`
}

func Benchmark_ExprStmt(b *testing.B) {
	os.Chdir("") // want `os\.Chdir\(\) could be replaced by b\.Chdir\(\) in .+`
}

func Test_ExprStmt(t *testing.T) {
	os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_AssignStmt(t *testing.T) {
	err := os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
}

func Test_AssignStmt_ignore_return(t *testing.T) {
	_ = os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_IfStmt(t *testing.T) {
	if err := os.Chdir(""); err != nil { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		// foo
	}
}

func TestName_RangeStmt(t *testing.T) {
	for i := range 5 {
		os.Chdir(strconv.Itoa(i)) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_ForStmt(t *testing.T) {
	for i := 0; i < 3; i++ {
		os.Chdir(strconv.Itoa(i)) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_DeferStmt(t *testing.T) {
	defer os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr(t *testing.T) {
	t.Log(os.Chdir("")) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr_deep(t *testing.T) {
	t.Log(
		fmt.Sprintf("here: %s, %s",
			strings.TrimSuffix(
				strings.TrimPrefix(
					fmt.Sprintf("%s",
						os.Chdir(""), // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
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
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}()
}

func Test_GoStmt_arg(t *testing.T) {
	go func(err error) {}(os.Chdir("")) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_CallExpr_recursive(t *testing.T) {
	foo(t, "")
}

func foo(t *testing.T, s string) error {
	return foo(t, fmt.Sprintf("%s", os.Chdir(s))) // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_FuncLit_ExprStmt(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		})
	}
}

func Test_SwitchStmt(t *testing.T) {
	switch {
	case runtime.GOOS == "linux":
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_SwitchStmt_case(t *testing.T) {
	switch {
	case os.Chdir("") == nil: // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
		// noop
	}
}

func Test_DeclStmt(t *testing.T) {
	var err error = os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
}

func Test_DeclStmt_tuple(t *testing.T) {
	var err, r error = errors.New(""), os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	_ = err
	_ = r
}

func Test_SelectStmt(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
			}
		}
	}()
}

func Test_DeferStmt_wrap(t *testing.T) {
	defer func() {
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}()
}

func Test_SelectStmt_anon_func(t *testing.T) {
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-doneCh:
				func() {
					os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
				}()
			}
		}
	}()
}

func Test_BlockStmt(t *testing.T) {
	{
		os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	}
}

func Test_TypeSwitchStmt(t *testing.T) {
	os.Chdir("") // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
}

func Test_TypeSwitchStmt_AssignStmt(t *testing.T) {
	switch v := os.Chdir("").(type) { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	case error:
		_ = v
	}
}

func Test_SwitchStmt_Tag(t *testing.T) {
	switch os.Chdir("") { // want `os\.Chdir\(\) could be replaced by t\.Chdir\(\) in .+`
	case errors.New(""):
	}
}

func foobar() {
	os.Chdir("")
}
