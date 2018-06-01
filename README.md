# go-plugin-sandbox

```
$ make -B
go build -buildmode=plugin -o hello/greeter.so hello/greeter.go
go build -buildmode=plugin -o bye/greeter.so bye/greeter.go
go run main.go
Hello
go run main.go -m bye
Bye
```
