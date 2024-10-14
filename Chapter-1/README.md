# Chapter 1
Setting up and hello world!
(Still a work in progress)

## Install the tools
The first step is to download and install the [Go programming language](https://go.dev/dl/)

You will also need a text editor. [VSCode](https://code.visualstudio.com/) is standard these days.

## Starting a new Go project
In the terminal, make a new go module called `site`:

```
go mod init site
```

This creates a file declaring what go version we are using.
Once we install modules, this file will also track our imported modules and their versions.


Now, create a new file called `main.go` with the following code:
```
package main

import "fmt"

func main() {
	fmt.Println("Hello shell!")
}

```

To make a webserver we can add this code to `main.go`.


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


Why the router works the way it does:
https://go.dev/blog/routing-enhancements 

Documentation for the router URL things
https://pkg.go.dev/net/http@go1.23.2#hdr-Patterns-ServeMux
(Why are these docs so hard to find?)
(Go docs are great if you know what you are looking for. But very approachable for beginners)
