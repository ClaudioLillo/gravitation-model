build:
	go build -o ./dist ./cmd/main.go

run:
	go run ./cmd/main.go

tidy:
	go mod tidy

play:
	ffplay output.avi

reset:
	rm -rf images && mkdir images
