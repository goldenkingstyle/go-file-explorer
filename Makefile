build:
	cls
	@go build -o ./bin/main.exe

run: build
	@./bin/main.exe

test:
	@go test -v