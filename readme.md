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

## Checklist

- [x] Route
- [ ] Handlers
  - [x] Single
  - [x] Resources (index,create,store,show,edit,update,destroy)
  - [x] Form Request Validation
  - [ ] Request Facade
    - [x] Validator()
    - [x] CurrentUrl()
    - [x] FullUrl()
    - [x] PreviousUrl()
    - [x] RedirectPrevious()
    - [x] SetFlash()
    - [x] GetFlash()
    - [x] SetFlashMap()
    - [x] GetFlashMap()
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
