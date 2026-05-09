# FakeAPI

Inspired by [json-server](https://github.com/typicode/json-server), but in Go.  
A fake REST server that generates CRUD endpoints from a JSON file. Data is kept in memory.

## Usage

```bash
# Create a data file (e.g. db.json)
echo '{
  "authors": [
    { "id": 1, "name": "anthony", "age": 21 },
    { "id": 2, "name": "miguel", "age": 32 }
  ],
  "books": [
    { "id": 1, "title": "Go 101", "author": "anthony" }
  ]
}' > db.json

go run cmd/main.go -f db.json
# Server started {"port": 8080}
```

Each key in the JSON becomes a REST endpoint under `/api/`.

## Endpoints

### GET /api/{resource}

List all records.

```bash
curl http://localhost:8080/api/authors
```

| Query   | Type | Default | Description |
|---------|------|---------|-------------|
| `limit` | int  | 10      | Max records |
| `offset`| int  | 0       | Start index |
| `sort`  | string | "key" | Sort field |

```bash
curl "http://localhost:8080/api/authors?limit=5&offset=0&sort=age"
```

### GET /api/{resource}/{id}

Get a record by its 1-based index.

```bash
curl http://localhost:8080/api/authors/1
```

### POST /api/{resource}

Add a new record.

```bash
curl -X POST http://localhost:8080/api/authors \
  -H "Content-Type: application/json" \
  -d '{"name": "jorge", "age": 28}'
```

### PUT /api/{resource}/{id}

Replace a record entirely.

```bash
curl -X PUT http://localhost:8080/api/authors/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "anthony l.", "age": 22}'
```

### PATCH /api/{resource}/{id}

Partially update a record (merges only fields that exist in the destination).

```bash
curl -X PATCH http://localhost:8080/api/authors/1 \
  -H "Content-Type: application/json" \
  -d '{"age": 25}'
```

### DELETE /api/{resource}/{id}

Delete a record.

```bash
curl -X DELETE http://localhost:8080/api/authors/1
```

## Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--port` | `-p` | `8080` | Server port |
| `--file` | `-f` | `fakeapi.json` | JSON data file |

## Run with the bundled example

```bash
go run cmd/main.go -f internal/parser/example.json
curl http://localhost:8080/api/authors
```

## Install

```bash
go install github.com/AnthonyLonsMax/fakeapi@v1.1.3
```

## Tests

```bash
go test -v -count 1 ./...
```
