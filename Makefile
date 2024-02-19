build:
	@go build -o myapp main.go

run: build
	@./myapp