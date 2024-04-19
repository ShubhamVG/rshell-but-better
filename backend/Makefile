gobuild:
	go build -o ./build/ ./cmd/server
	go build -o ./build/ ./cmd/client

launch-server: gobuild
	./build/server

clean:
	rm ./build/*
