# ogen-demo

A small Go HTTP service built step by step to learn [ogen](https://github.com/ogen-go/ogen),
the tool that generates Go server code from an OpenAPI v3 spec.

## Why this repo exists

BidOne's Go services (such as `api-orders` and `api-products`)
are built with ogen. Reading them is hard at first: it is not obvious which code is
generated and which is hand-written, or what a change to `openapi.yaml` will do to the Go code.

This repo rebuilds the same pattern from zero, one small step at a time, so that each
piece can be seen in isolation before meeting it in the real services.

The layout, layering and naming copy the BidOne repos on purpose, so what is learned
here transfers directly.

## Learning goals

One goal per lesson:

1. **ogen basics** - write a spec, generate a Go server from it, and tell generated
   code apart from hand-written code.
2. **BidOne layout and layering** - structure the service like the BidOne repos
   (handler -> service -> repo, pack/unpack) and map spec fields to Go types
   (`OptString`, `OptNilString`, ...).
3. **Errors** - return the right status code for each failure (such as 404), and keep
   internal Go errors out of the response body.
4. **ogen configuration** - understand what `config.yaml` turns on or off in the
   generated code, and which settings BidOne relies on.
5. **Security** - protect endpoints with a token, and see what ogen checks and
   what is left to our own code.
6. **Client** - generate an ogen client and call one service from another, the way
   BidOne's upstream services call the downstream ones.
7. **Multi-file spec** - split a large `openapi.yaml` into several files and bundle
   them back into one before generating.
8. **BFF (backend for frontend)** - pass data between services that each have their own
   generated types for the same schema.


## How to run

Run these from the repo root in PowerShell.

Regenerate the ogen code after changing `api/openapi.yaml`:

```powershell
make ogen
```

Start the server on port 8080:

```powershell
go run ./cmd/server
```

Try it:

```powershell
Invoke-RestMethod http://localhost:8080/v1/products
```

```powershell
Invoke-RestMethod -Method Patch -Uri http://localhost:8080/v1/products/1 -ContentType 'application/json' -Body '{"price":4.5}'
```

The data lives in memory, so restarting the server resets it.

## Progress

Each step is one commit; the hash links a step to `git log`. Each line says the
problem first, then what the step did about it.

### Lesson 1 - ogen basics (done)

Uses a flat layout (`generate.go` at the root, code in `internal/oas`); Lesson 2 replaces it.

- **Step 1** `15a52c4` - ogen needs a contract to generate from: write a first
  `openapi.yaml` and a `//go:generate` line that runs ogen.
- **Step 2** `14bb2f4` (+ typo fix `d8ab92b`) - the ogen command is not installed:
  add it to `go.mod`.
- **Step 3** `e0778ad` - see what ogen produces: run `go generate ./...`; every file
  in `internal/oas` is generated, then `go mod tidy` adds the packages it imports.
- **Step 4** `a0f5a6e` - generated code only declares the `Handler` interface:
  write the product handler that implements it.
- **Step 5** `1985c9a` - nothing listens for requests yet: start an HTTP server in
  `cmd/server`.
- **Step 6** `6f6ad28` - how does a field that may be missing look in Go? Add an
  optional `description`; it becomes `OptString`.
- **Step 7** `3d06dcb` - what does a spec change cost? Add `listProducts`; the build
  breaks until the handler gets the new method.

### Lesson 2 - BidOne layout and layering (done)

- **Step 1** `7f67eae` - the flat layout does not look like the BidOne repos: move
  ogen into its own `tools/` module and generate into `internal/handler/ogen`,
  as `api-orders` does.
- **Step 2** `0c11ce8` - the generate and build commands are long to type: add
  `make ogen` and `make server`.
- **Step 3** `339b609` - one handler file does not scale to many domains: add
  `/healthz`, group operations with `x-ogen-operation-group`, one folder per domain,
  joined by `CombinedHandler`.
- **Step 4** `dbff7c5` - HTTP code, business rules and data access were mixed:
  split into handler -> service -> repo, and `pack` ogen types at the edge.
- **Step 5** `3814f58` - a partial update must tell "not sent", `null` and a value
  apart: add `PATCH /v1/products/{id}` with `OptString` / `OptNilString`,
  `unpack` and `pkg/optional`.

### Lesson 3 - Errors (done)

- **Step 1** `0a02ab0` - "product not found" answered 400 (= the request is
  malformed): change it to 404 in the spec; the handler code does not change.
- **Step 2** `d6c9904` - a second failure reason (the new name is already used)
  needs its own status: add 409. One `*Error` can no longer pick the status, so ogen
  generates `PatchProductNotFound` / `PatchProductConflict`; the service returns
  `ErrProductNameTaken` and the handler maps it.
- **Step 3** `67314bf` - what does a client get for an error we did not plan? Make
  id 13 fail like a database: the 500 body leaks an internal hostname, nothing is
  logged, and errors come in four formats.
- **Step 4** `08f99a7` - fix the leak and the missing log: a custom
  `WithErrorHandler` returns the reason for 4xx, and logs 5xx while returning only
  "Internal Server Error". The body uses the spec's `Error` schema.
- **Step 5** (no commit, skipped) - JSON bodies for unknown paths and methods via
  `WithNotFound` / `WithMethodNotAllowed`; trap: the latter also receives `OPTIONS`
  (CORS preflight). BidOne keeps the defaults.
- **Step 6** (no commit, experiment) - ogen's "convenient errors" (`NewError`,
  enabled when every operation has the same `default` response) splits errors into
  two paths and turns a missing status into 200. BidOne sets
  `convenient_errors: "off"`.
- **Step 7** `05226d3` - remove the fake database failure with `git revert` of Step 3.
