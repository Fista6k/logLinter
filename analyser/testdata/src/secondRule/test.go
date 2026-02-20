package secondrule

import "log/slog"

func test() {
	slog.Info("гикаю")             // want "log message should contain only english symbols"
	slog.Info("тяжелый свэг, бро") // want "log message should contain only english symbols"
}
