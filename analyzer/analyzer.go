package analyzer

import (
	"go/ast"
	"strconv"
	"strings"
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

		ast.Inspect(file, func(f ast.Node) bool {
			expr, ok := f.(*ast.CallExpr)
			if !ok {
				return true
			}
			if !isLog(expr) {
				return true
			}
			if len(expr.Args) == 0 {
				return true
			}

			Checker(expr.Args[0], pass)
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

	// log/slog
	if pack == "slog" || pack == "log" {
		logMethods := map[string]bool{
			"Println": true, "Printf": true, "Print": true,
			"Info": true, "Error": true, "Warn": true, "Debug": true,
		}
		return logMethods[funcName]
	}

	// zap
	zapMethods := map[string]bool{
		"Info": true, "Infow": true, "Infof": true,
		"Error": true, "Errorw": true, "Errorf": true,
		"Warn": true, "Warnw": true, "Warnf": true,
		"Debug": true, "Debugw": true, "Debugf": true,
	}
	return zapMethods[funcName]
}

// логика проверки по 4 правилам
func Checker(arg ast.Expr, pass *analysis.Pass) {

	basicLit, ok := arg.(*ast.BasicLit)
	if !ok {
		pass.Reportf(arg.Pos(), "log should not contain variables for safety")
		return
	}

	text, err := strconv.Unquote(basicLit.Value)
	if err != nil {
		return
	}

	// rule 1: лог должен начинаться с маленькой буквы
	trim := strings.TrimSpace(text)
	if len(trim) != 0 {
		firstChar := rune(trim[0])
		if engCheck(firstChar) && unicode.IsUpper(firstChar) {
			pass.Reportf(arg.Pos(), "log should start with lower case")

		}
	}

	// rule 2: лог должен быть только на английском языке
	for _, char := range text {
		if unicode.IsLetter(char) && !engCheck(char) {
			pass.Reportf(arg.Pos(), "log should be only in English")
			break
		}
	}
	// rule 3: лог не должен содержать спец. символы
	for _, char := range text {
		if !checkedIfAllowed(char) {
			pass.Reportf(arg.Pos(), "log should not contain symbols")
			break
		}
	}

}

func checkedIfAllowed(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsNumber(char) || unicode.IsSpace(char) {
		return true
	}

	return false
}
func engCheck(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
