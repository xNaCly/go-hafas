# go-hafas

> A Hafas Go implementation providing qol helpers and abstractions for easier hafas interaction

**WARNING**: go-hafas is still a work in progress

## Features

- Abstractions over:
  - Authentication
  - HTTP request managment
  - Response deserialization and unwrapping
- Typesafe hafas API interaction
- Exhaustive testing suite
- Built-in Language support
- Configurable abstraction via functional modifiers
- Full access to underlying generated HAFAS types and endpoints for advanced
  use and users (via `Client.ClientWithResponses`)

## Install

```shell
go get github.com/xNaCly/go-hafas
go mod tidy
```

## Examples

### Initialisation and Configuration

```go
package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	hafas "github.com/xnacly/go-hafas"
	"github.com/xnacly/go-hafas/language"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	client, err := hafas.NewClient(
		"http://localhost:1234",
		"f4640da4-7586-4615-9621-0372542e9225",
		hafas.WithLanguage(language.ES),
		hafas.WithContext(ctx),
		hafas.WithHttpClient(&httpClient),
	)

	if err != nil {
		slog.Error("hafas", "msg", "failed to create hafas client", "err", err)
	}

	err = client.Init()
	if err != nil {
		slog.Error("hafas", "msg", "failed to init hafas client", "err", err)
	}

	err = client.Ping()
	if err != nil {
		slog.Error("hafas", "msg", "failed to ping hafas remote", "err", err)
	}
}
```

### Searching for Locations

#### By Name

#### By Coordinates

### Arrivals

### Departures

### DataInfo

### TripSearch

### JourneyDetail

### JourneyPos

## Testing

> The testing requires the `BASEURL` and `AUTH` env variables to be set to
> valid hafas values, see [`.env`](./.env) and fill these

### With nix flakes

```shell
dev
go test ./... -v
```

### Without nix

```shell
source ./.env
go test ./... -v
```
