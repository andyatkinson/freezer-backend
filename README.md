## Freezer Backend

Backend service using [Golang Echo](https://echo.labstack.com/)

Currently this only works in a local development environment.


### Local Development

Running locally:

```
go run server.go
```

### Installing Dependencies

Example:

```
go get -u github.com/lib/pq/...
```


## API

#### Listing Items

```
$ curl -H 'Content-Type: application/json' localhost:1323/items
```

#### Creating Items

```
curl -d '{"name":"item","addedOn":"2020-01-01"}' -H 'Content-Type: application/json' localhost:1323/items
{"name":"item","addedOn":"2020-01-01"}
```