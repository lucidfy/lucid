# GORVEL (UNDER DEVELOPMENT...)

Gorvel, yet another framework inspired with Laravel / Symfony structure, but written in Go!

## Documentation

... still in progress

## Auto Compile

Serving go requires a tedeous way of re-compiling codes, but there's a solution to automatically compile any `.go` files everytime you made changes. We're using reflex to hot reload our `cmd/serve.go`, here's how to install it.

```bash
$> go install github.com/cespare/reflex@latest
```

## Run via Docker

Just execute `./setup-docker` and it should create a docker image called **gorvel-img** and it will build a container **gorvel-container**

The docker container will consume these ports 8080 for gorvel and 8081 for svelte-kit, however these ports were internal, it forwards back to your docker host under these ports 8080 -> **8330**, 8081 -> **8331**. You can verify this by running `docker ps -a`

Therefore, try to open your browser and access http://localhost:8330 for gorvel and http://localhost:8331 for svelte-kit

## TODO's

- [x] Routing
- [x] Middlewares
- [x] CSRF Protection
- [x] Handlers (a.k.a "Controllers")
  - [x] Single
  - [x] Resources (index,create,store,show,edit,update,destroy)
  - [x] Form Request Validation
- [ ] Request
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
  - [x] GetUserAgent()
  - [x] GetFileByName()
  - [x] GetFiles()
- [ ] URL
  - [x] CurrentUrl()
  - [x] FullUrl()
  - [x] PreviousUrl()
  - [x] RedirectPrevious()
- [ ] Response
  - [x] Json()
  - [x] View()
  - [x] ViewWithStatus()
- [ ] Session
  - [x] SetFlash()
  - [x] GetFlash()
  - [x] SetFlashMap()
  - [x] GetFlashMap()
- [ ] Console
  - [x] route:defined
  - [x] route:registered
  - [ ] make:resource
    - [ ] Model
    - [ ] C.R.U.D. Handlers
- [x] Pagination
- [x] ORM using gORM (by default)
- [x] Storage
  - [x] Get()
  - [x] Put()
  - [x] Delete()
  - [x] Exists()
  - [x] Missing()
  - [x] Size()
  - [x] Path()

## Security Issues

Please sent a direct email to me for any vulnerable things you may find via: daison12006013@gmail.com
