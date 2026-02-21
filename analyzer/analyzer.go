package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var logFunctions = map[string]bool{
	"Info":   true,
	"Error":  true,
	"Warn":   true,
	"Debug":  true,
	"Infof":  true,
	"Errorf": true,
	"Warnf":  true,
	"Debugf": true,
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "logLinter",
		Doc:      "Check are logs right",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if !isLogFunc(pass, call) {
			return
		}

		if len(call.Args) == 0 {
			return
		}

		message := extractMessage(call.Args[0])
		if message == "" {
			return
		}

		checkFirstLetter(pass, call, message)

		checkEnglish(pass, call, message)

		checkSpecialChars(pass, call, message)

		checkSensitive(pass, call, message)
	})

	return nil, nil
}

func isLogFunc(pass *analysis.Pass, call *ast.CallExpr) bool {
	var funcName string
	var sel *ast.SelectorExpr
	var ok bool

	if sel, ok = call.Fun.(*ast.SelectorExpr); !ok {
		return false
	}
	funcName = sel.Sel.Name

	if !logFunctions[funcName] {
		return false
	}

	if obj, ok := pass.TypesInfo.ObjectOf(sel.X.(*ast.Ident)).(*types.PkgName); ok {
		if obj.Imported().Path() == "log/slog" {
			return true
		}
	}

	switch t := pass.TypesInfo.TypeOf(sel.X).(type) {
	case *types.Pointer:
		if named, ok := t.Elem().(*types.Named); ok {
			if named.Obj().Pkg() != nil && named.Obj().Pkg().Path() == "go.uber.org/zap" {
				return true
			}
		}
	}

	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "log" {
			return true
		}
	}

	return false
}

func extractMessage(arg ast.Expr) string {
	switch lit := arg.(type) {
	case *ast.BasicLit:
		if lit.Kind == token.STRING {
			return strings.Trim(lit.Value, "\"")
		}
	case *ast.BinaryExpr:
		return extractMessage(lit.X)
	}

	return ""
}
