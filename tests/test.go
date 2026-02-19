package tests

import (
	"log/slog"

	"go.uber.org/zap"
)

func testLogs() {
	slog.Error("Little error1")
	slog.Info("Something here1")

	slog.Error("little error1")
	slog.Info("something here1")
}

func testZap(logger *zap.Logger) {
	logger.Info("Something here2")
	logger.Error("Little error2")

	logger.Info("something here2")
	logger.Error("little error2")
}
