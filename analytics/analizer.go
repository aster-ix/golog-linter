package analizer 

import("golang.org/x/tools/go/analysis")

var Analyzer = &analysis.Analyzer{
	Name: "golog-linter",
	Doc:  "linter for golang logs",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	
}