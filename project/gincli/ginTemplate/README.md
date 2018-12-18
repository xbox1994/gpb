# Gin API Demo
## How to run

In one terminal, run the server:

```bash
$ go run main.go
```

In another terminal, try making requests:

```bash
curl -i http://localhost:8080/users
```

```bash
curl -i http://localhost:8080/users/1
```

```bash
curl -X POST \
  http://localhost:8080/users \
  -H 'content-type: application/json' \
  -d '{"id": 1, "name":"wty"}'
```