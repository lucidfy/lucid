# Configuration

- [# .env](#-env)
- [# .env.local](#-envlocal)

---

Lucid is heavily relying to `.env` files and they're automatically loaded when calling `os.Getenv()`

To set an env on runtime, you just need to call `os.Setenv("MY_KEY", "VALUE")`

{#-env}

## [#](#-env) .env

```bash
APP_ENV=local
```

The `.env` file will be loaded first, lucid parses and checks the `APP_ENV`, assume the value is equal to **"local"**, therefore it will load the `.env.local`

> Note: The main github repository of Lucid uses "local", consider not to commit your `.env` by adding it inside your `.gitignore`

<br>

```bash
APP_TIMEZONE="UTC"
```

By default we're using UTC time, this is useful when storing data to be in-sync with our application, to check which timezone to use, refer to "TZ database name" at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

<br>

```bash
APP_LANGUAGE="en-US"
```

By default we're using `"en-US"` stands for "English United States" translation, it will automatically use `resources/translations/en_US.go`

<br>

```bash
SESSION_NAME="lucid_session"
SESSION_DOMAIN=""
SESSION_LIFETIME=7200
```

Every http request, we're producing a session identifier and that is stored inside a guest's browser, learn more about it [here](/session)

<br>

```bash
CONSOLE_PATH=/app/commands
HANDLERS_PATH=/app/handlers
MIDDLEWARES_PATH=/app/middlewares
MODELS_PATH=/app/models
DATABASE_PATH=/databases
TRANSLATION_PATH=/resources/translations
VIEW_PATH=/resources/views/go
ROUTES_PATH=/routes
STORAGE_PATH=/storage
```

Here we are pointing our folder structure, mainly used for generating files, ***see example below:***

```go
path.Load().BasePath("path/to/folder/file.pdf") // /Users/johndoe/my-lucid-project/path/to/folder/file.pdf
path.Load().ConsolePath("") // .../app/commands
path.Load().HandlersPath("") // .../app/handlers
path.Load().MiddlewaresPath("") // .../app/middlewares
path.Load().ModelsPath("")  // .../app/models
path.Load().DatabasePath("") // .../databases
path.Load().TranslationPath("") // .../resources/translations
path.Load().ViewPath("") // .../resources/views/go
path.Load().RoutesPath("") // .../routes
path.Load().StoragePath("") // .../storage
path.Load().SessionPath("") // .../storage/framework/sessions
```

{#-envlocal}

## [#](#-envlocal) .env.local

```bash
APP_DEBUG=true
```

This is mainly used to trace any unwanted http request errors.

<br>

```bash
APP_KEY="SomeRandomString"
```

The above is used as our key to encrypt / decrypt a string, this is mainly used when generating user's cookie, generating passwords and so on.

> To learn more about the [Lucid's Cryptography](/cryptography)

<br>

```bash
LOGGING_ENABLED=true
LOGGING_FILE=/storage/logs/lucid.log
```

This enables or disables our logger, as of the moment our logger is based on filesystem, but later on we will add a driver based, such as kibana logs and so on!

<br>

```bash
SCHEME="http"
HOST="0.0.0.0"
PORT="8080"
```

This configures how lucid will serve your app

<br>

```bash
CSRF_AUTH_KEY="MyCsrfAuthKey"
CSRF_TRUSTED_ORIGIN=""
```

This will prevent any [Cross Site Request Forgery](https://owasp.org/www-community/attacks/csrf), provide the authentication key and your trusted origin.

<br>

```bash
DB_CONNECTION="sqlite"
DB_DATABASE="sqlite.db"
```

Here we set the database & connection, also supports well known RDBMS `mysql`, `postgres`, `sqlserver` and `clickhouse`

```bash
DB_CONNECTION="mysql"
DB_DATABASE="user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

DB_CONNECTION="postgres"
DB_DATABASE="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

DB_CONNECTION="sqlserver"
DB_DATABASE="sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"

DB_CONNECTION="clickhouse"
DB_DATABASE="tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
```
