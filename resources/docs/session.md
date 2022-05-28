# Session

- [# Basics](#-basics)
- [# Adaptors](#-adaptors)
- [# API](#-api)
  - [# Get](#-get)
  - [# Set](#-set)
  - [# Destroy](#-destroy)
  - [# SetFlash](#-setflash)
  - [# GetFlash](#-getflash)
  - [# SetFlashMap](#-setflashmap)
  - [# GetFlashMap](#-getflashmap)

---

Session is mainly used to store other information of the current state, we often use this to set the user's language, user's preferred themes and whatever things that we can store.

> Note: Most developers often compare `Session` over `Cache`, they have the same process, its just that `Cache` is a global and it is not bound to any identifier, while `Session` is bound to `lucid_session`  cookie as identifier.

{#-basics}

## [#](#-basics) Basics

> Note: this is a continuation guide from [Middleware -> Authenticated Middleware](/middleware#-authenticated-middleware)

Please refer to the file at `handlers/auth_handler/auth.go` as a good example

```go
// handlers/auth_handler/auth.go @ LoginAttempt
ses := engine.Session
ses.Set("authenticated", record.ID)
...

// handlers/auth_handler/auth.go @ User
ses := engine.Session
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

{#-api}

## [#](#-api) API

Here are the lists of functions available with their sample

{#-get}

### [#](#-get) Get

```go
ses.Get("key")
```

{#-set}

### [#](#-set) Set

```go
ses.Set("key", 100)
```

{#-destroy}

### [#](#-destroy) Destroy

```go
ses.Destroy("key")
```

{#-setflash}

### [#](#-setflash) SetFlash

```go
ses.SetFlash("message", "Hello Jane!")
```

{#-getflash}

### [#](#-getflash) GetFlash

```go
ses.GetFlash("message")
```

{#-setflashmap}

### [#](#-setflashmap) SetFlashMap

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

### [#](#-getflashmap) GetFlashMap

```go
ses.GetFlashMap("messages")
```
