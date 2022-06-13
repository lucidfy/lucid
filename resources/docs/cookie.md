# Cookie

- [# Basics](#-basics)
- [# API](#-api)
  - [# CreateSessionCookie](#-createsessioncookie)
  - [# Set](#-set)
  - [# Get](#-get)
  - [# Expire](#-expire)

---

Each time a guest visits your website, we're producing a session identifier and that is stored inside the guest's browser, this is made possible inside our `app/middlewares/session.go`, by inspecting your browser network, there should be a `lucid_session`

Succeeding requests should carry that cookie, the identifier's value is a random string and encrypted using your `APP_KEY`

{#-basics}

## [#](#-basics) Basics

As an example for this is located at `app/middleares/session.go`

```go
// app/middleares/session.go
coo := cookie.New(w, r)
_, err := coo.Get(os.Getenv("SESSION_NAME"))

if err != nil && errors.Is(err, http.ErrNoCookie) {
    coo.CreateSessionCookie()
}
```

The code above, we're checking if there is existing cookie from the http request header, if the error contains `http.ErrNoCookie`, then we need to call `CreateSessionCookie()` to generate the `lucid_session`.

{#-api}

## [#](#-api) API

Here are the lists of functions available with their sample

{#-createsessioncookie}

### [#](#-createsessioncookie) CreateSessionCookie

```go
coo.CreateSessionCookie()
```

{#-set}

### [#](#-set) Set

```go
coo.Set("user_id", 100)
```

{#-get}

### [#](#-get) Get

```go
coo.Get("user_id")
```

{#-expire}

### [#](#-expire) Expire

```go
coo.Expire("user_id")
```
