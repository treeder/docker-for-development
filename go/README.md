## Developing Go (golang) using Docker

This example will show you how to compile your code with the same architecture we have on IronWorker so it will
run properly.

**NOTE**: It's probably best to copy this code into your GOPATH and work from there. Go is hard to work with using Docker
alone so you'll have to install Go to your system. 

Let's build hello.go and run it. First, we need to install our dependencies. 

```sh
go get github.com/iron-io/iron_go/worker
```

Build it:

```sh
docker run --rm -v "$GOPATH":/gopath -v "$(pwd)":/app -w /app google/golang sh -c 'go build -o hello'
```

Notice we mounted our local GOPATH into the container. 

Now run it:

```sh
docker run --rm -v "$(pwd)":/app -w /app google/golang sh -c './hello'
```

Boom.
