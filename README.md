cli
===

This is the home of the Eval.gd CLI.

The program is written in Go and uses the Eval.gd JSON api for everything it
does.

See manpage.md for application details.


## Building

Building `evalgd` requires a working Golang environment. Run `go run evalgd.go`
to run the program, or to just build it, run `go build evalgd.go`, which will
result in a binary called `evalgd`.

The manpage is built with [Ronn](https://github.com/rtomayko/ronn) and should
be kept up to date. To build the manpage, simply run `ronn man/evalgd.md`.

Ronn is installable as a Rubygem.
