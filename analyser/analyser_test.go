package logLinter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLinter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, NewAnalyzer(), "firstRule")
	analysistest.Run(t, testdata, NewAnalyzer(), "secondRule")
	//analysistest.Run(t, testdata, NewAnalyzer(), "thirdRule")
	analysistest.Run(t, testdata, NewAnalyzer(), "fourthRule")
}
