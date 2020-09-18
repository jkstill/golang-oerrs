
Oracle Errors Package for golang
================================

Be gentle in you critcism, this is my first go package...

Using all of the codes from OH/rdbms/mesg/oraus.msg, generate a package to lookup and compare error codes

Several Errors can be looked up by a common name

All of the others can be looked up by number

## Get and use oerrs

```text
   go get github.com/jkstill/golang-oerrs
```

As per comments in the program, 'go run' and 'go build' take ~1 minute to complete, and so does 'go get'.

This seems to be due to 20k+ calls to `errors.New`.

Once built, the lookup time is quite fast.

If already installed, and you want the most recent version:

```text
   go get -u github.com/jkstill/golang-oerrs
```

Simple test program, using the built in test function

```go
package main

import (
   "github.com/jkstill/golang-oerrs"
)

func main () {
   oerrs.Test()
}
```

## Generate oerrs.go

This is just a bash script to generate oerrs.go

Commenting out 'mode=BUILD' near the top of the script will generate a go program that can be run on the CLI

```text
  ./gen-oracodes-for-golang.sh  > oerrs.go
  go run oerrs.go
```

As per comments in the program, 'go run' and 'go build' take ~1 minute to complete.

This seems to be due to 20k+ calls to `errors.New`.

Once built, the lookup time is quite fast.
 
