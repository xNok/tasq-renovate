
build:
	CGO_ENABLED=0 go build -o taskq_renovate ./cmd/main.go
	docker build -t taskq-renovate .
