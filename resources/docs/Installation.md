# Installation

## Manual Setup

I should assume you've successfully installed your go in your machine, to start working with this, you should fork a copy of the `develop` branch your github.

If you want to quickly try Lucid, please follow bellow source, make sure your port `8080` is open to serve your local http, or modify your lucid `.env` file.

```bash
$> echo $GOPATH
/Users/johndoe/go
$> cd /Users/johndoe/go
$> mkdir src/
$> git clone git@github.com:YOURUSERNAME/lucid.git
$> cd lucid
$> go mod download
$> bash ./serve
```

## Docker Setup

Make sure you have installed [Docker Desktop](https://www.docker.com/products/docker-desktop/), after that run `./setup-docker` under your lucid's root folder

```sh
~/go/src/lucid: ./setup-docker
```

After that, once you've successfully build your docker, to `tail -f` the console output

```sh
docker logs -f lucid-container
```

To **stop** or **start** again the docker container

```sh
docker container stop lucid-container
docker container start lucid-container
```
