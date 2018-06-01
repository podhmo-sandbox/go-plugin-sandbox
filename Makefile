default: build
	go run main.go
	go run main.go -m bye

build: hello/greeter.so bye/greeter.so
hello/greeter.so: hello/*.go
	go build -buildmode=plugin -o $@ $^
bye/greeter.so: bye/*.go
	go build -buildmode=plugin -o $@ $^
