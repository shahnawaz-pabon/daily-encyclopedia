## Up and Run

Go to the `single-stage` directory, replace that `Dockerfile` by this directory's `Dockerfile` and run the following commands.

```sh
docker build -t go-multi-stage .
docker run -it -p 8080:8080 go-multi-stage
```

## Image size

If you run

```sh
sudo docker images
```

You'll see the following output

```sh
REPOSITORY       TAG                 IMAGE ID       CREATED              SIZE
go-multi-stage   latest              890df974b404   About a minute ago   12.1MB
```

## Comparison

This was the size of `single-stage`'s docker image

```sh
REPOSITORY        TAG                 IMAGE ID       CREATED         SIZE
go-single-stage   latest              1373d418fb49   8 seconds ago   337MB
```
