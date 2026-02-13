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
			expr, ok := f.(*ast.CallExpr)
			if !ok {
				return true
			}

			testansw := isLog(expr)
			fmt.Println(testansw, " = islog")
			return true
		})
	}

	return nil, nil
}

func isLog(expr *ast.CallExpr) bool {
	selectorExpr, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// funcPack := selectorExpr.Sel.Name -- имя пакета

	funcPack, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	pack := funcPack.Name
	fmt.Println(pack)

	if pack == "slog" || pack == "log" {
		return true
	}
	return false
}
