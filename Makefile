# hello:
# 	echo "Hello"

# build:
# 	go build -o bin/main main.go

run:
	LOCATION=$(location) go run cmd/cluster-agent/main.go