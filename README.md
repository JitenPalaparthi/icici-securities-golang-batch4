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

### Keywords(22 of 25)

```
break        default      func         interface    
case         defer                     map          struct
             else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Builtin (18)
