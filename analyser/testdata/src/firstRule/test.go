package firstrule

import (
	"log/slog"
)

func test() {
	slog.Error("Little error1")  // want "log message should start with lowercase letter"
	slog.Info("Something here1") // want "log message should start with lowercase letter"
}
