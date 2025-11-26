# user-service
API HTTP + gRPC para manejo de usuarios — proyecto portfolio.

## Requisitos
- Go 1.20+
- Postgres (opcional para usar postgres repository)
- protoc (opcional para regenerar `.pb.go`)

## Rápido
```bash
cp .env.example .env
make build
make run
# o para grpc
go run cmd/grpc/main.go
```

Tests:
```
make test
```

Generar protos (si tienes `protoc`):
```
protoc --go_out=grpc/gen --go-grpc_out=grpc/gen grpc/proto/user.proto
```

Badge: (ficticio) Build: passing
