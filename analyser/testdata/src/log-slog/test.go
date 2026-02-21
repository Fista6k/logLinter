package logslog

import "log/slog"

func test() {
	msg := "Error"

	//first rule
	slog.Info(msg)               // OK
	slog.Error("Little error1")  // want "log message should start with lowercase letter"
	slog.Info("Something here1") // want "log message should start with lowercase letter"

	slog.Info(" error happened")
	slog.Info("error")

	//second rule
	slog.Info("гикаю")             // want "log message should contain only english symbols"
	slog.Info("тяжелый свэг, бро") // want "log message should contain only english symbols"
	slog.Info("server started!🚀")  // want "log message should contain only english symbols"
	slog.Debug("№;%*№(())")        // want "log message should contain only english symbols"

	slog.Error("only english") // OK

	//third rule
	slog.Error("connection failed!!!") // want "log message shouldn't contain any '!!!'"
	slog.Error("connection failed...") // want "log message shouldn't contain any '...'"

	slog.Info("here we go again!!") // OK
	slog.Info("here we go again..") // OK

	//fourth rule
	password := "12345"
	slog.Info("user password: " + password)             // want "log shouldn't contain any sensitive words"
	slog.Error("user password: " + password + "broooo") // want "log shouldn't contain any sensitive words"
	slog.Info("user password: 12345")                   // want "log shouldn't contain any sensitive words"

	slog.Info("user:") // OK
}
