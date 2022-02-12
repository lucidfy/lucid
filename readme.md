# GORVEL (UNDER DEVELOPMENT!)

Gorvel is just a boilerplate inspired under the foldering structure from Laravel's client, although the framework lives inside the `pkg` and `internal`.

## Executable Dependencies

- We're using reflex to hot reload our `serve.go`, here's how it works
  - Everytime you change a `.go` code that lives inside this project, it will autocompile.

```bash
$> go install github.com/cespare/reflex@latest
```

## How to?

I should assume you've installed your go in your machine, to start working with this, you make a clone under your `$GOPATH/src/`

```bash
$> echo $GOPATH
/Users/johndoe/go

$> cd /Users/johndoe/go
$> mkdir src/
$> wget -c https://github.com/daison12006013/gorvel/archive/refs/heads/master.tar.gz -O - | tar -xz
$> cd gorvel-master
$> bash ./serve
```
