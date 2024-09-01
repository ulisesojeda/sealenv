.PHONY: build

build:
	env GOOS=linux GOARCH=amd64 go build -o build/sealenv_linux_amd64 main.go
	env GOOS=linux GOARCH=arm64 go build -o build/sealenv_linux_arm64 main.go
	env GOOS=darwin GOARCH=amd64 go build -o build/sealenv_darwin_amd64 main.go
	env GOOS=darwin GOARCH=arm64 go build -o build/sealenv_darwin_arm64 main.go
	env GOOS=windows GOARCH=amd64 go build -o build/sealenv_windows_amd64.exe main.go
