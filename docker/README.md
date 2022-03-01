<div align="center">
  <a href="https://www.docker.com/">
    <img alt="docker" src="../logos/docker.png"/>
  </a>
  <h1>Docker</h1>
</div>

# Table of Contents

- [Miscellaneous commands](#miscellaneous-commands)

  - [See docker networks](#see-docker-networks)
  - [Link ngrok to the running docker container](#link-ngrok-to-the-running-docker-container)

## Miscellaneous commands

### See docker networks

```sh
# This command lists all the docker networks. docker-compose will create a new network when you run docker-compose up
docker network ls
```

### Link ngrok to the running docker container

```sh
# The --rm flag tells docker that the container should automatically be removed after we close docker.
# The -it flag tells docker that it should open an interactive container instance, -it is short for --interactive + --tty. When you docker run with this command it takes you straight inside the container.
docker run --rm -it --link <container-name> --net <docker-network> shkoliar/ngrok http <container-name>:<running-port> --authtoken <auth-token-from-ngrok-dashboard>
```

**NOTE**

- You can see container list by running `docker container ls`
- Get your authtoken from the [ngrok's dashboard](https://dashboard.ngrok.com/login).
