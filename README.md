GoLang
======

My _super awesome_ Go code!

Here you will find some code I wrote while learning Google's [Go-Lang programming language](http://golang.org).

So far, I have some code that can create zip archives and extract them. I'm pretty sure I'll have more soon.

Building and Running Programs
-----------------------------

To run a program quickly, use `go run`, for example:

```sh
cd go-by-example/values
go run values.go
```

To build and run a program, for example, the _hello-world go-by-example_ program, use the following command:

```sh
cd go-by-example/hello-world
go build -o hello-world.exe hello-world.go
./hello-world.exe
```

Formatting
----------

To automatically format go source files, use `go fmt`:

```sh
cd go-by-example/variables
go fmt variables.go
```

This will modify the source in-place if it does not meet the source code style guide from the community.

References
----------

* [go.dev](https://go.dev/doc/tutorial/getting-started)
* [Go by Example](https://gobyexample.com/)
* [Go 101](https://go101.org/article/101.html)
