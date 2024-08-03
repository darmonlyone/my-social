# Social

A speed run application for Go language API

## Requirement

!!require Go

- Go
- sqlboiler
- golang-migrate
- (postgres server)
  - If don't have but your own. Can deploy docker on `deployment/postgres`.
  - require **docker**
  - run `make infra-up`

If you’ve never used SQLBoiler, download the code-gen binaries:

```
$ go install github.com/volatiletech/sqlboiler/v4@latest
$ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```

If you’ve never used golang-migrate

```
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Structure

![structure](img/structure.png)

A simple structure for my social app. Have only Account table and Post table for this application

## How to run

Require `make` if necessary.

- `make run` run the project
- `make test` run test

- `make infra-up` start the postgres server if don't have it own postgres server

if continue implement

- `make migrate-up` for migrate the data
- (need if update database structure) `make gensqlboiler` gen boiler for the project

### !! Require

Setup environment `.env` and `sqlboiler.toml`. You can look up on `.example` suffix file.

## API Doc

<!-- TODO add api doc -->
