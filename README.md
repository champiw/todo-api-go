# Todo API

A small RESTful Todo API built with Go.

This is primarily a learning project.

## Development

This project uses Nix flakes for its development environment.

```bash
nix develop
go run ./cmd/server
```

The API currently runs on:

```text
http://localhost:8000
```

## Project Structure

```text
todo-api/
├── cmd/
│   └── server/
│       └── main.go
├── flake.nix
├── flake.lock
├── go.mod
└── go.sum
```
