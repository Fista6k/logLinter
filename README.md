# Log-Linter
Custom Go analyzer for checking logs in go projects, integrate with golangci-lint (go plugin .so) Works with Linux/MacOS or wsl on Windows

Features
* Checking is first letter of log lowercase
* Checking is log message containing only english letters
* Checking is log message containing spec symbols
* Checking log message for sensitive data

Project Structure
```
Linter/
|--analyzer/
|  |--testdata/src
|  |  |--go.uber.org/zap
|  |  |  |--zap.go
|  |  |--log-slog/
|  |  |  |--test.go
|  |  |--zap/
|  |  |  |--test.go
|  |--analyzer.go
|  |--analyzer_test.go
|  |--rules.go
|--cmd/logLinter
|  |--main.go
|--plugin/
|  |--plugin.go
|--go.mod
|--go.sum
|--README.md
```

Setup
1. Clone repository
  ```
  git clone https://github.com/Fista6k/logLinter
  ```

2. Build the plugin (wsl or linux)
  ```
  cd logLinter
  make build
  ```

3. Copy new logLinter.so file into your project

4. Add .golangci.yml
  ```
  version: "2"

  linters:
    enable:
      - loglinter

  settings:
    custom:
      loglinter:
        path: ./logLinter.so
        description: "Check logs"
        original-url: github.com/Fista6k/logLinter
  ```

5. Run the linter (in your project directory) wsl or linux
  ```
  golangci-lint run
  ```

Main problem
- You need to have the same golangci-lint, go and go/tools versions to completely start it on your project
