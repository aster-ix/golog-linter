package analyzer

import (
	"go/ast"
	"go/token"
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

			Checker(text, pass, arg.Pos())
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

	if pack == "slog" || pack == "log" {

		//TODO:вот эту страшную проверку сделать через мап потом,если времени хватит
		if funcName == "Println" || funcName == "Printf" || funcName == "Info" ||
			funcName == "Error" || funcName == "Warn" || funcName == "Debug" {
			return true
		}
	}
	return false
}

func Checker(text string, pass *analysis.Pass, pos token.Pos) bool {

	firstChar := rune(text[0])
	if unicode.IsUpper(firstChar) {
		pass.Reportf(pos, "- log should start with lower case")
	}

	for _, char := range text {
		if unicode.IsLetter(char) && !engCheck(char) {
			pass.Reportf(pos, "- log should be only in English")
			break
		}
	}

	for _, char := range text {
		if !checkedIfAllowed(char) {
			pass.Reportf(pos, "- log should not contain symbols")
			break
		}
	}

	return true
}

func checkedIfAllowed(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsNumber(char) || unicode.IsSpace(char) {
		return true
	}
	return false
}
func engCheck(char rune) bool {
	return (char >= 'a' && char <= 'r') || (char >= 'A' && char <= 'Z')
}
