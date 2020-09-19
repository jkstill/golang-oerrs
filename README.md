
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

## oracle-err-catch.go

This is a small program to demonstrate connecting to an Oracle database and running a query.

If the table does not exist, it will be created.

```text
jkstill:~/go/src/oracle/oracle-test$ ./oracle-err-catch 
That Table does not exist
Table mytable created - please try again

jkstill:~/go/src/oracle/oracle-test$ ./oracle-err-catch 
ID:    1 The date is: 2020-09-18T21:17:48-07:00
ID:    2 The date is: 2020-09-18T21:17:48-07:00
2020/09/18 21:17:48 id= 2
```
 
## Named Errors

Following are the errors that currently may be looked up by name:

- uniqueConstraint
	- ORA-00001: unique constraint (%s.%s) violated")
- invalidIdentifier
	- ORA-00904: %s: invalid identifier")
- tableNotFound
	- ORA-00942: table or view does not exist")
- badConnectID
	- ORA-12154: TNS:could not resolve the connect identifier specified")
- badConnectPath
	- ORA-12198: TNS:could not find path to destination")
- cannotConnect
	- ORA-12203: TNS:unable to connect to destination")
- listenerFailed
	- ORA-12500: TNS:listener failed to start a dedicated server process")
- badAddress
	- ORA-12533: TNS:illegal ADDRESS parameters")
- noSuchDB
	- ORA-12545: Connect failed because target host or object does not exist")
- lostContact
	- ORA-12547: TNS:lost contact")
- connectionClosed
	- ORA-12537: TNS:connection closed")
- storedProcErr
	- ORA-04063: %s has errors")
- storedProcInvalid
	- ORA-04064: not executed, invalidated %s")
- TNSAdapterErr
	- ORA-12560: TNS:protocol adapter error")
- snapshotTooOld
	- ORA-01555: snapshot too old: rollback segment number %s with name \"%s\" too small")
- invalidCredentials
	- ORA-01017: invalid username/password; logon denied")
- missingExpression
	- ORA-00936: missing expression")
- invalidNumber
	- ORA-01722: invalid number")
- PL/SQLValueErr
	- ORA-06502: PL/SQL: numeric or value error%s")
- invalidCharacter
	- ORA-00911: invalid character")
- invalidSQL
	- ORA-00933: SQL command not properly ended")
- bug
	- ORA-00600- internal error code, arguments: [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s]")
- invalidNumber
	- ORA-01722: invalid number")
- EOFOnChannel
	- ORA-03113: end-of-file on communication channel")
- maxCursors
	- ORA-01000: maximum open cursors exceeded")
- coreDump
	- ORA-07445: exception encountered: core dump [%s] [%s] [%s] [%s] [%s] [%s]")

