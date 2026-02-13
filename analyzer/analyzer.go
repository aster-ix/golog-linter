package analyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gologlinter",
	Doc:  "linter for golang logs",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(f ast.Node) bool {
			expr, err := f.(*ast.CallExpr)
			if !err {
				return true
			}

			fmt.Println(expr, " == expr")
			return true
		})
	}

	return nil, nil
}
