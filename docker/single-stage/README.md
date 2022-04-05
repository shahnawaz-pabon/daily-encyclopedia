## Manually Up and Run

```sh
go run main.go
```

## Via docker

```sh
docker build -t go-single-stage .
docker run -it -p 8080:8080 go-single-stage
```

Open up [http://localhost:8080](http://localhost:8080) in your browser and you'll be able to see `Home Page` message which is returned from the Go application.

If you open up `index.html` in your browser, you'll see these two messages on the browser console.

```sh
Attempting Connection...
Successfully Connected
```

and `Hi From the Client!` message on the running terminal.

## Image size

If you run

```sh
sudo docker images
```

You'll see the following output

```sh
REPOSITORY        TAG                 IMAGE ID       CREATED         SIZE
go-single-stage   latest              1373d418fb49   8 seconds ago   337MB
```

## Comparison

This was the size of `multi-stage`'s docker image

```sh
REPOSITORY       TAG                 IMAGE ID       CREATED              SIZE
go-multi-stage   latest              890df974b404   About a minute ago   12.1MB
```
