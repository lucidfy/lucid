# Deployment

Deploying lucid requires a little bit of configuration, especially handling crashes; there's a good online tutorial on how to achieve this using **Supervisord**

> https://www.socketloop.com/tutorials/how-to-automagically-restart-your-crashed-golang-server

Now we need to build our lucid into binary

## ./build

Running `./build` will create a go bin files inside `.bin/` folder.

```bash
# ./build
LUCID_ROOT=${PWD} ./cmd/build-go.sh
```

### ./cmd/build-go.sh

```bash
# windows
GOOS=windows GOARCH=amd64 go build -o ./.bin/lucid-windows.exe ./cmd/serve/main.go

# darwin (mac os)
GOOS=darwin GOARCH=amd64 go build -o ./.bin/lucid-darwin ./cmd/serve/main.go

# linux
GOOS=linux GOARCH=amd64 go build -o ./.bin/lucid-linux ./cmd/serve/main.go

# default
go build -o ./.bin/lucid ./cmd/serve/main.go
```
