cli
===

This is the home of the Eval.so CLI.

The program is written in Go and uses the Eval.so JSON api for everything it
does.

See manpage.md for application details.


## Building

Building `evalso` requires a working Golang environment. Run `go run evalso.go`
to run the program, or to just build it, run `go build evalso.go`, which will
result in a binary called `evalso`.

The manpage is built with [Ronn](https://github.com/rtomayko/ronn) and should
be kept up to date. To build the manpage, simply run `ronn man/evalso.md`.

Ronn is installable as a Rubygem.
