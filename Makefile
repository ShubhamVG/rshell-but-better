gobuild:
	go build -o ./build/ ./cmd/server
	go build -o ./build/ ./cmd/client

run: gobuild
	./build/server

clean:
	rm ./build/*