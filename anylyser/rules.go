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
	if containEmoji(message) {
		pass.Reportf(call.Pos(), "log shouldn't contain any emojis: %q", message)
	}

	if strings.Contains(message, "...") {
		pass.Reportf(call.Pos(), "log message shouldn't contain any '...': %q", message)
	}

	if strings.Contains(message, "!!!") {
		pass.Reportf(call.Pos(), "log message shouldn't contain any '!!!': %q", message)
	}

	if containAnyPuncSymbols(message) {
		pass.Reportf(call.Pos(), "log shouldn't contain any punkt symbols except '.' and ','", message)
	}
}

func containEmoji(msg string) bool {
	for _, r := range msg {
		if unicode.Is(unicode.So, r) {
			return true
		}
	}
	return false
}

func containAnyPuncSymbols(msg string) bool {
	for _, r := range msg {
		if unicode.IsPunct(r) && r != '.' && r != ',' {
			return true
		}
	}
	return false
}

func checkSensetive(pass *analysis.Pass, call *ast.CallExpr, message string) {

}
