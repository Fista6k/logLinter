package logLinter

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var sensitiveWords = []string{
	"password",
	"api_key",
	"apikey",
	"token",
	"secret",
	"auth",
}

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
	if strings.Contains(message, "...") {
		pass.Reportf(call.Pos(), "log message shouldn't contain any '...': %q", message)
		return
	}

	if strings.Contains(message, "!!!") {
		pass.Reportf(call.Pos(), "log message shouldn't contain any '!!!': %q", message)
		return
	}
}

func checkSensitive(pass *analysis.Pass, call *ast.CallExpr, message string) {
	lowerMsg := strings.ToLower(message)
	if strings.ContainsAny(lowerMsg, ":=") {
		for _, w := range sensitiveWords {
			if strings.Contains(lowerMsg, w) {
				pass.Reportf(call.Pos(), "log shouldn't contain any sensitive words: %q", message)
				return
			}
		}
	}
}
