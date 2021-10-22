build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/server ./cmd/server