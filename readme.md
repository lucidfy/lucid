# GORVEL (UNDER DEVELOPMENT...)

Gorvel, yet another framework inspired with Laravel / Symfony structure, but written in Go!

## Documentation

... still in progress

## Auto Compile

Serving go requires a tedeous way of re-compiling codes, but there's a solution to automatically compile any `.go` files everytime you made changes. We're using reflex to hot reload our `cmd/serve.go`, here's how to install it.

```bash
$> go install github.com/cespare/reflex@latest
```

## How to Contribute?

I should assume you've successfully installed your go in your machine, to start working with this, you should fork a copy of `master` branch to your github, checkout under your `$GOPATH/src/` folder.

If you want to quickly try Gorvel, please follow bellow source, make sure your port `8080` is open to serve your local http.

```bash
$> echo $GOPATH
/Users/johndoe/go
$> cd /Users/johndoe/go
$> mkdir src/
$> wget -c https://github.com/daison12006013/gorvel/archive/refs/heads/master.tar.gz -O - | tar -xz
$> cd gorvel-master
$> bash ./serve
```

## Deploy as a Docker Container

... still in progress

## TODO's

- [x] Route
- [ ] Handlers
  - [x] Single
  - [x] Resources (index,create,store,show,edit,update,destroy)
  - [x] Form Request Validation
  - [ ] Url Facade
    - [x] CurrentUrl()
    - [x] FullUrl()
    - [x] PreviousUrl()
    - [x] RedirectPrevious()
  - [ ] Session Facade
    - [x] SetFlash()
    - [x] GetFlash()
    - [x] SetFlashMap()
    - [x] GetFlashMap()
  - [ ] Request Facade
    - [x] Validator()
    - [x] All()
    - [x] Get()
    - [x] GetFirst()
    - [x] Input()
    - [x] HasContentType()
    - [x] HasAccept()
    - [x] IsForm()
    - [x] IsJson()
    - [x] IsMultipart()
    - [x] WantsJson()
    - [x] GetIp()
  - [ ] Response Facade
    - [x] Json()
    - [x] View()
    - [x] ViewWithStatus()
- [ ] Console Commands
  - [x] route:defined
  - [x] route:registered
  - [ ] make:resource
    - [ ] Model
    - [ ] C.R.U.D. Handlers
- [x] Pagination
- [x] ORM using gORM (by default)

## Security Issues

Please sent a direct email to me for any vulnerable things you may find via: daison12006013@gmail.com
