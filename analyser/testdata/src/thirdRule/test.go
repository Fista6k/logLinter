package thirdrule

import "log/slog"

func test() {
	slog.Info("server started!🚀")      // want "log message should contain only english symbols"
	slog.Error("connection failed!!!") // want "log message shouldn't contain any '!!!'"
}
