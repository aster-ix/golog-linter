package main

import (
	"github.com/aster-ix/golog-linter/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

type AnalyzerPlugin struct{}

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
