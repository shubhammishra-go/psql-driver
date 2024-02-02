# pSQL Driver in GO
it is a pure Go Postgres driver for the database/sql package. 

for more documentation refer this...

[link](https://pkg.go.dev/github.com/lib/pq)

## To Use
--> first intialize a repository [Optional]
<br>
--> install Psql package

## Intialize

```go mod init```
or specified a name to your repository like this
```go mod init github.com/shubhammishra-1```

## Install Psql driver package

```bash
go get github.com/lib/pq
```

## import it into your main file

``` _ go get github.com/lib/pq```

## Connection String

to connect with psql database a string like URL must be parse

```postgres://postgres:55120@localhost/?sslmode=disable```

or Connecting with database directly

```postgres://postgres:55120@localhost/database_name?sslmode=disable```



After setuping see the code how to perform SQL operations




## Psql libraries for Go
sqlX
```https://github.com/jmoiron/sqlx```

pgX

```https://github.com/jackc/pgx```

