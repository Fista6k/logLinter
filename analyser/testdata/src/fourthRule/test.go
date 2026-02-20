package fourthrule

import "log/slog"

func test() {
	password := "12345"
	slog.Info("user password: " + password) // want "log shouldn't contain any sensitive words"
}
