# Docker Installation {#\" class="text-4xl"}

Make sure you have installed [Docker Desktop](https://www.docker.com/products/docker-desktop/), after that run `./setup-docker` under your gorvel's root folder

```sh
~/go/src/gorvel: ./setup-docker
```

After that, once you've successfully build your docker, to `tail -f` the console output

```sh
docker logs -f gorvel-container
```

To **stop** or **start** again the docker container

```sh
docker container stop gorvel-container
docker container start gorvel-container
```
