# Installation

## Manual and only Lucid Go

I should assume you've successfully installed your go in your machine, to start working with this, you should fork a copy of the `develop` branch your github.

If you want to quickly try Lucid, please follow bellow source, make sure your port `8080` is open to serve your local http, or modify your lucid `.env` file.

```bash
> wget -c https://github.com/lucidfy/lucid/archive/refs/heads/develop.tar.gz -O - | tar -xz
> cd lucid-develop
> go mod download
> bash ./serve
```

## Docker Setup with Svelte

Make sure you have installed [Docker Desktop](https://www.docker.com/products/docker-desktop/), after that run `./setup-docker` under your lucid's root folder

```
> wget -c https://github.com/lucidfy/setup/archive/refs/heads/develop.tar.gz -O - | tar -xz
> mv setup-develop lucid-setup
> cd lucid-setup/
> wget -c https://github.com/lucidfy/lucid/archive/refs/heads/develop.tar.gz -O - | tar -xz
> mv lucid-develop/ src/lucid/
> wget -c https://github.com/lucidfy/ui/archive/refs/heads/develop.tar.gz -O - | tar -xz
> mv ui-develop/ src/lucid-ui/
> ./setup-container
```

After that, once you've successfully built your own container, to tail the console stdout for both go and sveltekit, try below command.

```sh
> docker logs -f lucid-container
```

To **stop** or **start** again the docker container

```sh
> docker container stop lucid-container
> docker container start lucid-container
```
