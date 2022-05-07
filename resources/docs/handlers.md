# Handlers

A handler process an http request, it will only be available after the middleware had been iterated.

## Single Handler

Here's how to generate a single handler

```bash
./run make:handler healthcheck
Created a handler at app/handlers/healthcheck.go
```

## Resource Handler

To generate a resource handler

```bash
./run make:resource reports
Created a resource handlers at app/handlers/reportshandler/
```
