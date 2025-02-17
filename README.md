# icici-securities-golang-batch4
Golang Training for ICICI


### Kind of projects

1. Binary/Exec 
2. Library/Package

## Keywords (25)

- func, package (2 of 25)

## Builtin Functions(18)

- print, println

## Go commands

- Go mod

```
go mod init demo
```

- Run 

```
go run main.go
go run greet.go main.go
go run .
```

## Go build

``` 
go build .
go build -o hello .
```

## Go env

1. GOROOT 
2. GOPATH
3. GOBIN

4. GOOS
5. GOARCH

- TO know cross compilation options

```
go tool dist list
```

- Cross compile and build 

```
GOOS=linux GOARCH=amd64 go build -o build/demo-linux-amd64 .
```

- release mode builod 

```
go build -ldflags="-w -s" -o build/demo-mapc-arm64-small .
```