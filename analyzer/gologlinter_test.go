package analyzer_test

import (
	"path/filepath"
	"testing"

	gologlinter "github.com/aster-ix/golog-linter/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata, err := filepath.Abs(filepath.Join("..", "testdata"))
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, gologlinter.Analyzer, "tests.go")
}
