# go-mega-tutorial
A multi page tutorial for building fullstack websites in Go

Chapter 1: Go hello worlds (Move basic HTML here?)
Chapter 2: HTML, forms and Database (Standard lib) (Making a comment section/wall)
Chapter 3: Architectue & tooling (Introduce the idea that these tools really speed up development) Architectue
Chapter 4: Templ, CSS
Chapter 5: HTMX
Chapter 6: Databases with SQLC
Chapter 7: Databases migrations with Goose (Combine with SQLC?)
Chapter 8: Auth/login
Chapter 9: Logging and error handling
Chapter 10: Deployment to prod (fly.io/Docker and VPS)

Later:
CSS Frameworks and/or modern CSS?
API services like email/sms
Include how to use Git somewhere? Maybe not.



## Chapter 1 notes
- Install Go, VSCode (Or your perfered editor)
- Making a go project
    go mod init site
- Hello world! in the shell
- Hello world! in the browser

Go mod (As generated)
```
module site

go 1.22.7

```

Hello shell
```
package main

import "fmt"

func main() {
	fmt.Println("Hello shell!")
}

```

Hello web
```
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello shell!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", root)
	http.ListenAndServe(":4000", mux)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello web!")
}
```
(TODO explain the GET, {$} and w,r vars and types)



## Future
Use sqlc and goose? https://pressly.github.io/goose/blog/2024/goose-sqlc/ (How do bools work with sqlite?)
This is a Go way of doing thigs. But with tools to speed things up

Change to the Go Fullstack Tutorial? Or include some API things?
Write a corresponding blog post with design decisions for the tutorial. (This is also a great place to get feedback)
