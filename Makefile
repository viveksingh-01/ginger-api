#build
build:
	@go build -o bin/ginger-api cmd/main.go

#run
run: build
	@./bin/ginger-api