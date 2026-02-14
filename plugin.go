package gologlinter

import (
	"github.com/aster-ix/golog-linter/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("gologlinter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return &PluginGolog{}, nil
}

type PluginGolog struct{}

func (p *PluginGolog) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *PluginGolog) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
