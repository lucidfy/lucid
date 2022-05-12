# Handlers

A handler process an http request, it will only be available after the middleware had been iterated.

## Single Handler

Here's how to generate a single handler

```bash
./run make:handler healthcheck

Created handler, located at:
> ~/lucid/app/handlers/healthcheck.go
```

## Resource Handler

To generate a resource handler

```bash
./run make:resource reports

Created resource handler, located at:
 > ~/lucid/app/handlers/reports_handler/update.go
 > ~/lucid/app/handlers/reports_handler/create.go
 > ~/lucid/app/handlers/reports_handler/delete.go
 > ~/lucid/app/handlers/reports_handler/lists.go

Created model, located at:
 > ~/lucid/app/models/reports/model_test.go
 > ~/lucid/app/models/reports/model.go
 > ~/lucid/app/models/reports/struct.go

Created validation, located at:
 > ~/lucid/app/validations/reports.go

```
