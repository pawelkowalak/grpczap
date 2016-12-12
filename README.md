# gRPCZapLogger

Small `grpclog.Logger` interface implementation to allow using [zap logger](https://github.com/uber-go/zap) with gRPC.

Typical usage:

```go
package main

import (
    "github.com/uber-go/zap"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

var logger zap.Logger

func init() {
    logger = zap.New(zap.NewTextEncoder(zap.TextTimeFormat(time.RFC3339)), zap.Output(os.Stdout))
    grpclog.SetLogger(grpczap.New(logger))
}

func main() {
    conn, err := grpc.Dial("127.0.0.1:31337", grpc.WithInsecure())
    if err != nil {
        logger.Error("Can't open gRPC connection", zap.Error(err))
        return
    }
}
```
