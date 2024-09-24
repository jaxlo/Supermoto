# go-mega-tutorial
A multi page tutorial for building fullstack websites in Go


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
