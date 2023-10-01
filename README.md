# Simple Golang API Example
## What it is?
This project is a simple Go API that simulates the basic functions of a library management system. It allows users to see what books are available, check out books, return books, and add new books to the library's collection.

## How it works
- Check `go.mod` for requirements.
- Run `go run main.go` to start the server on port 8080 (default) or you can change this by editing line 107.
- Then either use postman or simple curl comments to get results

## Simple curls commands
Get request for **/books** using curl

    curl localhost/8080/books

Get request for **/books/:id** using curl

    curl localhost:8080/books/2

Post request for **/books** using curl 

    curl localhost:8080/books --header "Content-Type: application/json" -d @bookExample.json --request "POST"

Patch request for **checkout** using curl

    curl localhost:8080/checkout?id=1 --request "PATCH"

Patch request for **/return** using curl

    curl localhost:8080/return?id=1 --request "PATCH"

## Notes
- ID of the books may be similar. It is because books are kept in array. No real database operations implemented, yet.