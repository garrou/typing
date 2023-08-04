# typing

## Docker Compose

```sh
docker compose up -d
```

## Docker

```sh
docker run --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=typing -d postgres:15.3-alpine3.18
```