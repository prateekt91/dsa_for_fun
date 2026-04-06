# DSA for Fun

A hands-on collection of data structure and algorithm implementations built from scratch, purely for learning and exploration.

## Implementations

### Go

| Topic | File | Description |
|-------|------|-------------|
| Hash Map | [`go/hashmap_impl.go`](go/hashmap_impl.go) | Custom hash map with separate chaining for collision resolution. Supports `Put`, `Get`, and `Remove` operations using a polynomial rolling hash (`hash * 31 + byte`). |

## Running

**Prerequisites:** [Go](https://go.dev/dl/) installed on your machine.

```bash
go run go/hashmap_impl.go
```

## Project Structure

```
dsa_for_fun/
├── Readme.md
└── go/
    └── hashmap_impl.go
```

