// Package main contains the basic runnable version of the linter.
package main

import (
	"github.com/ldez/usetesting"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(usetesting.NewAnalyzer())
}
