package analyzer

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gologlinter",
	Doc:  "linter for golang logs",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {

		// TODO: удалить личные комменты потом
		ast.Inspect(file, func(f ast.Node) bool {
			expr, ok := f.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isLog(expr) {
				return true
			}

			arg := expr.Args[0]
			basicLit, ok := arg.(*ast.BasicLit)
			if !ok {
				return true
			}

			// text := basicLit.Value -- string - выводится с кавычками
			text, err := strconv.Unquote(basicLit.Value)
			if err != nil {
				return true
			}

			if !Checker(text) {
				fmt.Println("ошибка")
			} else {
				fmt.Println("нет ошибки")
			}

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

	funcName := selectorExpr.Sel.Name
	funcPack, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	pack := funcPack.Name
	fmt.Println(pack)

	if pack == "slog" || pack == "log" {

		//TODO:вот эту страшную проверку сделать через мап потом,если времени хватит
		if funcName == "Println" || funcName == "Printf" || funcName == "Info" ||
			funcName == "Error" || funcName == "Warn" || funcName == "Debug" {
			return true
		}
	}
	return false
}

func Checker(text string) bool {

	firstChar := rune(text[0])
	if unicode.IsUpper(firstChar) {
		return false
	}

	return true
}
