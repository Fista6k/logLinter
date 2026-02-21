package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLinter(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := NewAnalyzer()
	analysistest.Run(t, testdata, analyzer, "log-slog")
	analysistest.Run(t, testdata, analyzer, "zap")
}
