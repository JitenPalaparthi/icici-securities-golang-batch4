# icici-securities-golang-batch4

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

- release mode build

```
go build -ldflags="-w -s" -o build/demo-mapc-arm64-small .
```

- To fetch all dependencies

```
go mod tidy
```

- To vendor dependencies in a vendor directory

```
go mod vendor
```

- go get a specific repository

```
go get github.com/jitenpalaparthi/somerepo
```

### Keywords(25 of 25)
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Builtin (18 of 18)
- append, close, clear, copy, print,println, cap,len, complex, real, imag, max,min, panic, recover, new, make, delete

### To run test

```
go test ./...
```
```
go test -timeout 30s -run ^TestStringReverse$ demo/models
```

```
go test -timeout 30s -run ^TestValidateContact$ demo/models
go test -timeout 30s -run ^(TestValidateContact|TestStringReverse1)$ demo/models
```

- Package level test
```
go test demo/models
```

- cover profile

```
go test --coverprofile=cover.out demo/models
```

- benchmark

```
go test -benchmem -run=^$ -coverprofile=bench.out -bench . demo/models
```

go test -benchmem -run=^$ -coverprofile=/var/folders/rs/kkl0cpr51t574sdxt7hc7fth0000gn/T/vscode-goMAgt1D/go-code-cover -bench . demo/models