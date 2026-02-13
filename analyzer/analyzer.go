package analyzer

import (
	"fmt"
	"go/ast"
	"strconv"

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

			fmt.Println(text)
			// Checker(text)
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
	//TODO: проверку на конкретную функцию, тк там есть log.New() который по факту лог, но текст не выводит
	return false
}

func Checker(text string) bool {

	// проверка на большую букву
	// letters := []rune(text)
	// if len(letters)>0{
	// 	firstChar := letters[0]
	// 	if unicode.IsUpper(firstChar) || firstChar
	// }

	return true
}
