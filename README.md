# comms
Connect and select

### Installation

comms requires latest [Golang](https://golang.org/doc/install) to run.

### Build the app (Optional)

```sh
$ cd comms
$ go build -o comms
$ ./comms -h
```

### Run the app (Not built)
Go to project folder and type:

```sh
$ cd comms
$ go run main.go
```

Optional parameters available
```sh
$ go run main.go -h
```

Wait for the program to load
```sh
$ go run main.go
$ Data loaded and processed in 28.391580286s
```

Start typing program ids separated by space
```sh
$ go run main.go
$ Data loaded and processed in 28.391580286s
$ 87 21 0 300
```

### Tests
Run the tests:
```sh
$ go test -cover ./... -v
```
### Documentation
Run the docs:
```sh
$ godoc -http=":6060"
```
Then visit: [http://localhost:6060/pkg/github.com/clybs/](http://localhost:6060/pkg/github.com/clybs/)