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

Each step is one commit. A new section is added when a lesson is finished.

### Lesson 1 - ogen basics (done)

- Step 1-3: write a first spec, install ogen, generate code with `go generate`.
- Step 4-5: implement a product handler and start an HTTP server.
- Step 6: add an optional field and see it become `OptString`.
- Step 7: add a second endpoint (list all products).

### Lesson 2 - BidOne layout and layering (done)

- Step 1: move the ogen tool into `tools/` and generate into `internal/handler/ogen`,
  the same way as `api-orders`.
- Step 2: add `make ogen` and `make server`.
- Step 3: add `/healthz`, group operations with `x-ogen-operation-group`,
  and join handlers with `CombinedHandler`.
- Step 4: split the code into handler, service and repo layers, with `pack`.
- Step 5: add `PATCH /v1/products/{id}`: `unpack`, `pkg/optional`, and the
  `OptString` / `OptNilString` types that tell "not sent", `null` and a value apart.
