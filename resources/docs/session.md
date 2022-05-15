# Session

- [# Basics](#-basics)
- [# Adaptors](#-adaptors)
  - [# Given Functions](#-given-functions)
    - [# Get](#-get)
    - [# Set](#-set)
    - [# Destroy](#-destroy)
    - [# SetFlash](#-setflash)
    - [# GetFlash](#-getflash)
    - [# SetFlashMap](#-setflashmap)
    - [# GetFlashMap](#-getflashmap)

---

Each time a guest visits your website, we're producing a session identifier and that is stored inside the guest's browser, this is made possible inside our `app/middlewares/session.go`, by inspecting your browser network, there should be a `lucid_session`

Succeeding requests should carry that cookie, the identifier's value is a random string and encrypted using your `APP_KEY`

{#-basics}

## [#](#-basics) Basics

> Note: this is a continuation guide from [Middleware -> Authenticated Middleware](/middleware#-authenticated-middleware)

Please refer to the file at `handlers/auth_handler/auth.go` as a good example

```go
// handlers/auth_handler/auth.go @ LoginAttempt
ses := session.File(w, r)
ses.Set("authenticated", record.ID)
...

// handlers/auth_handler/auth.go @ User
ses := session.File(w, r)
userID, err := ses.Get("authenticated")
...
```

{#-adaptors}

## [#](#-adaptors) Adaptors

Adapter      | Implemented? | Example
-------------|--------------|--------------------------
File         | Yes          | `session.File(w, r)`
~~Cookie~~   | Not yet      | `session.Cookie(w, r)`
~~Database~~ | Not yet      | `session.Database(w, r)`
~~Redis~~    | Not yet      | `session.Redis(w, r)`

{#-given-functions}

### [#](#-given-functions) Given Functions

Here are the lists of functions available with their sample

{#-get}

#### [#](#-get) Get

```go
ses.Get("key")
```

{#-set}

#### [#](#-set) Set

```go
ses.Set("key", 100)
```

{#-destroy}

#### [#](#-destroy) Destroy

```go
ses.Destroy("key")
```

{#-setflash}

#### [#](#-setflash) SetFlash

```go
ses.SetFlash("message", "Hello Jane!")
```

{#-getflash}

#### [#](#-getflash) GetFlash

```go
ses.GetFlash("message")
```

{#-setflashmap}

#### [#](#-setflashmap) SetFlashMap

```go
ses.SetFlashMap(
    "messages",
    map[string]interface{}{
        "john": "Hello Jane!",
        "jane": "Hi John!",
    },
)
```

{#-getflashmap}

#### [#](#-getflashmap) GetFlashMap

```go
ses.GetFlashMap("messages")
```
