# go-requestcontext

This repo provides a small shim for tagging a request ID onto the request context, and then get access to that request ID in sub-contexts.

The impetus for this code is that the setting and use of the request ID context key may be in different projects, and it may be desirable to decouple the dependency between such projects. Since context keys are advisted to be set using a custom sentinal type, this repo provides the place for that singleton value to live.

## Usage

```go
import (
    "context"
    "net/http"

    requestcontext "github.com/willscott/go-requestcontext"
)

func main() {
    // Injection
    mux := http.NewServeMux()
    handler := requestcontext.Middleware(mux, "X-RequestID")
    ...
}

func sendSubrequest(ctx context.Context) {
    rid := requestcontext.IDFromContext(ctx)
    ...
}

```

## Documentation

See [godoc](https://godoc.org/github.com/willscott/go-requestcontext) for API documentation.


## Contributing

Contributions are welcome! This repository is part of the IPFS project and therefore governed by our [contributing guidelines](https://github.com/ipfs/community/blob/master/CONTRIBUTING.md).

## License

[SPDX-License-Identifier: Apache-2.0 OR MIT](LICENSE.md)