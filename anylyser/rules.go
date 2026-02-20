package logLinter

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func checkFirstLetter(pass *analysis.Pass, call *ast.CallExpr, message string) {
	runes := []rune(strings.TrimSpace(message))
	if len(runes) == 0 {
		return
	}

	first := runes[0]

	if unicode.IsLetter(first) && !unicode.IsLower(first) {
		pass.Reportf(call.Pos(), "log message should start with lowercase letter: %q", message)
	}
}

func checkEnglish(pass *analysis.Pass, call *ast.CallExpr, message string) {
	for _, r := range message {
		if r > unicode.MaxASCII {
			pass.Reportf(call.Pos(), "log message should contain only english symbols: %q", message)
			return
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, call *ast.CallExpr, message string) {
	//not implemented yet
}

func checkSensetive(pass *analysis.Pass, call *ast.CallExpr, message string) {
	//not implemented yet
}
