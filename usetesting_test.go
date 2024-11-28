package usetesting

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir string
	}{
		{dir: "basic"},
		{dir: "dot"},
		{dir: "nottestfiles"},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			analysistest.Run(t, analysistest.TestData(), Analyzer, test.dir)
		})
	}
}
