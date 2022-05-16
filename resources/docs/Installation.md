# Installation

- [# Host Machine](#-host-machine)
  - [# Lucid "GO"](#-lucid-go)
  - [# Lucid "SvelteKit"](#-lucid-sveltekit)
- [# Via Docker](#-via-docker)

---

{#-host-machine}

## [#](#-host-machine) Host Machine

{#-lucid-go}

### [#](#-lucid-go) Lucid "GO"

I should assume you've successfully [installed your go](https://go.dev/dl/) in your machine.

If you want to quickly try Lucid, please follow below source, make sure your port `8080` is open to serve your local http, or modify your lucid `.env` file.

```bash
> wget -c https://github.com/lucidfy/lucid/archive/refs/heads/develop.tar.gz -O - | tar -xz
> cd lucid-develop
> go mod download
> bash ./serve
```

{#-lucid-sveltekit}

### [#](#-lucid-sveltekit) Lucid "SvelteKit"

```bash
> wget -c https://github.com/lucidfy/ui/archive/refs/heads/develop.tar.gz -O - | tar -xz
> cd ui-develop/
> npm install
> ./make guest dev -- --host=0.0.0.0 --port=8081
```

After executing above, it should automatically sere `src/guest/` and opens a browser pointing to localhost:8081

```bash
> ./make docs dev -- --host=0.0.0.0 --port=8082
```

Then running above should serve the `src/docs/`, it should automatically open a browser pointing to localhost:8082

---

{#-via-docker}

## [#](#-via-docker) Via Docker

Make sure you have installed [Docker Desktop](https://www.docker.com/products/docker-desktop/) in your machine, then download the [lucidfy/setup](https://github.com/lucidfy/setup)

```bash
 > wget -c https://github.com/lucidfy/setup/archive/refs/heads/develop.tar.gz -O - | tar -xz
 > mv setup-develop lucid-setup
 > cd lucid-setup/
```

After placing the lucidfy/setup into your machine, we then need to download the [Lucid Svelte](https://github.com/lucidfy/ui) and [Lucid Go](https://github.com/lucidfy/lucid)

```bash
> wget -c https://github.com/lucidfy/lucid/archive/refs/heads/develop.tar.gz -O - | tar -xz
> mv lucid-develop/ src/lucid/
> wget -c https://github.com/lucidfy/ui/archive/refs/heads/develop.tar.gz -O - | tar -xz
> mv ui-develop/ src/lucid-ui/
```

after setting up all the folders, therefore execute the `./setup-container`

```bash
> ./setup-container
```

After that, once you've successfully built your own container, to tail the console stdout for both go and sveltekit, try below command.

```bash
> docker logs -f lucid-container
```

To **stop** or **start** again the docker container

```bash
> docker container stop lucid-container
> docker container start lucid-container
```
