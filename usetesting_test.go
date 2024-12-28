package usetesting

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir     string
		options map[string]string
	}{
		{dir: "oschdir/basic"},
		{dir: "oschdir/dot"},
		{dir: "oschdir/nottestfiles"},
		{dir: "oschdir/disable", options: map[string]string{"oschdir": "false"}},

		{dir: "contextbackground/basic"},
		{dir: "contextbackground/dot"},
		{dir: "contextbackground/nottestfiles"},
		{dir: "contextbackground/disable", options: map[string]string{"contextbackground": "false"}},

		{dir: "contexttodo/basic"},
		{dir: "contexttodo/dot"},
		{dir: "contexttodo/nottestfiles"},
		{dir: "contexttodo/disable", options: map[string]string{"contexttodo": "false"}},

		{dir: "osmkdirtemp/basic"},
		{dir: "osmkdirtemp/dot"},
		{dir: "osmkdirtemp/nottestfiles"},
		{dir: "osmkdirtemp/disable", options: map[string]string{"osmkdirtemp": "false"}},

		{dir: "ossetenv/basic", options: map[string]string{"ossetenv": "true"}},
		{dir: "ossetenv/dot", options: map[string]string{"ossetenv": "true"}},
		{dir: "ossetenv/nottestfiles", options: map[string]string{"ossetenv": "true"}},
		{dir: "ossetenv/disable", options: map[string]string{"ossetenv": "false"}},

		{dir: "ostempdir/basic", options: map[string]string{"ostempdir": "true"}},
		{dir: "ostempdir/dot", options: map[string]string{"ostempdir": "true"}},
		{dir: "ostempdir/nottestfiles", options: map[string]string{"ostempdir": "true"}},
		{dir: "ostempdir/disable"},

		{dir: "oscreatetemp/basic"},
		{dir: "oscreatetemp/dot"},
		{dir: "oscreatetemp/nottestfiles"},
		{dir: "oscreatetemp/disable", options: map[string]string{"oscreatetemp": "false"}},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			t.Parallel()

			newAnalyzer := NewAnalyzer()

			for k, v := range test.options {
				err := newAnalyzer.Flags.Set(k, v)
				if err != nil {
					t.Fatal(err)
				}
			}

			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), newAnalyzer, test.dir)
		})
	}
}
