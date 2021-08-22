# CRUD API IN GO and POSTGRESQL 

## Running Project
1. Install GO
2. Install POSTGRESQL 
3. Export following env variables
```
export DIALECT="postgres"
export HOST="<DB URL>"
export DBPORT="=<DB PORT>"
export DBUSER="<DB USERNAME>"
export DBNAME="books"
export PASSWORD="<DB PASSWORD>"
```
4. Run following command
```
go build
go mod tidy
go run api.go
```

## API Documentation

1. Fetch All books
```
curl -XGET 'localhost:8764/books'
```

2. Fetch A book with id = 1
```
curl -XGET 'localhost:8764/book/1'
```

3. Add a Book
```
curl -XPOST -H "Content-type: application/json" -d '{
    "id": 1,
    "title": "Rich Dad Poor Dad",
    "author": "Sharon Lechter "
}' 'localhost:8764/book'
```

4. Update a Book with
```
curl -XPUT -H "Content-type: application/json" -d '{
    "id": 1,
    "title": "Rich Dad Poor Dad",
    "author": "Robert Kiyoski"
}' 'localhost:8764/book'
```

5. Delete a Book
```
curl -XDELETE -H "Content-type: application/json" -d '{
"id": 1,
}' 'localhost:8764/book'
```

