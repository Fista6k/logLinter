package zap

import "go.uber.org/zap"

func test(logger *zap.Logger) {
	msg := "Error"

	//first rule
	logger.Info(msg)               // OK
	logger.Error("Little error1")  // want "log message should start with lowercase letter"
	logger.Info("Something here1") // want "log message should start with lowercase letter"

	logger.Info(" error happened") // OK
	logger.Info("error")           // OK

	//second rule
	logger.Info("гикаю")             // want "log message should contain only english symbols"
	logger.Info("тяжелый свэг, бро") // want "log message should contain only english symbols"
	logger.Info("server started!🚀")  // want "log message should contain only english symbols"
	logger.Debug("№;%*№(())")        // want "log message should contain only english symbols"

	logger.Error("only english") // OK

	//third rule
	logger.Error("connection failed!!!") // want "log message shouldn't contain any '!!!'"
	logger.Error("connection failed...") // want "log message shouldn't contain any '...'"

	logger.Info("here we go again!!") // OK
	logger.Info("here we go again..") // OK

	//fourth rule
	password := "12345"
	logger.Info("user password: " + password)             // want "log shouldn't contain any sensitive words"
	logger.Error("user password: " + password + "broooo") // want "log shouldn't contain any sensitive words"
	logger.Info("user password: 12345")                   // want "log shouldn't contain any sensitive words"

	logger.Info("user:") // OK
}
