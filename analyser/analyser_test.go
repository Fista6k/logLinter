package logLinter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLinter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, NewAnalyzer(), "log-slog")
	analysistest.Run(t, testdata, NewAnalyzer(), "zap")
}
