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

### Data Seeding

* `create database freezer_development`
* Gorm is responsible for migrating the database

* The following seed data can be manually inserted. Gorm manages various additional columns about objects.

```
insert into items (name, added_on) VALUES ('Lasagna for 4 (4)', '2020-02-01');
insert into items (name, added_on) VALUES ('Frozen peas, 1 bag (8 oz.)', '2020-03-01');
insert into items (name, added_on) VALUES ('Popsicles (box) (8)', '2020-04-01');
```


## API

#### Listing Items

```
$ curl -H 'Content-Type: application/json' localhost:1323/items
```

#### Creating Items

As a JSON payload

```
curl -d '{"name":"Chicken Nuggets","addedOn":"2020-01-01"}' -H 'Content-Type: application/json' localhost:1323/items
```

## Heroku

Set Go version

```
heroku config:set GOVERSION=go1.9
```
