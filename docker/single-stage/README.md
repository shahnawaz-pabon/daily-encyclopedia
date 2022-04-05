## Manually Up and Run

```sh
go run main.go
```

## Via docker

```sh
docker build -t go-websocket .
docker run -it -p 8080:8080 go-websocket
```

Open up [http://localhost:8080](http://localhost:8080) in your browser and you'll be able to see `Home Page` message which is returned from the Go application.

If you open up `index.html` in your browser, you'll see these two messages on the browser console.

```sh
Attempting Connection...
Successfully Connected
```

and `Hi From the Client!` message on the running terminal.
