package main

import (
	logLinter "github.com/Fista6k/analyser"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	analyzer := logLinter.NewAnalyzer()
	singlechecker.Main(analyzer)
}
