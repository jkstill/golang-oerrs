
package oerrs

import (
	"errors"
	"fmt" 
)

/*

It takes a very long time, 1 minute, for go to compile the map ErrByNum.

There are 20k+ entries, and I thought perhaps the issue was memory allocation.

That is, each new entrie was requiring the allocation of memory.

So, why not tell go the size of the map?

so this:

ErrByName = make(map[string]error)
ErrByNum = make(map[int]error)

was changed to this:

ErrByName = make(map[string]error,100)
ErrByNum = make(map[int]error,22000)

There was no difference in compilation time, either via 'go run' or 'go build'


*/

var ErrByName map[string]error
var ErrByNum map[int]error

func init() {

	ErrByName = make(map[string]error,100)
	ErrByNum = make(map[int]error,22000)

	ErrByName["uniqueConstraint"] = errors.New("dpiStmt_execute: ORA-00001: unique constraint (%s.%s) violated")
	ErrByName["invalidIdentifier"] = errors.New("dpiStmt_execute: ORA-00904: %s: invalid identifier")
	ErrByName["tableNotFound"] = errors.New("dpiStmt_execute: ORA-00942: table or view does not exist")
	ErrByName["badConnectID"] = errors.New("dpiStmt_execute: ORA-12154: TNS:could not resolve the connect identifier specified")
	ErrByName["badConnectPath"] = errors.New("dpiStmt_execute: ORA-12198: TNS:could not find path to destination")
	ErrByName["cannotConnect"] = errors.New("dpiStmt_execute: ORA-12203: TNS:unable to connect to destination")
	ErrByName["listenerFailed"] = errors.New("dpiStmt_execute: ORA-12500: TNS:listener failed to start a dedicated server process")
	ErrByName["badAddress"] = errors.New("dpiStmt_execute: ORA-12533: TNS:illegal ADDRESS parameters")
	ErrByName["noSuchDB"] = errors.New("dpiStmt_execute: ORA-12545: Connect failed because target host or object does not exist")
	ErrByName["lostContact"] = errors.New("dpiStmt_execute: ORA-12547: TNS:lost contact")
	ErrByName["connectionClosed"] = errors.New("dpiStmt_execute: ORA-12537: TNS:connection closed")
	ErrByName["storedProcErr"] = errors.New("dpiStmt_execute: ORA-04063: %s has errors")
	ErrByName["storedProcInvalid"] = errors.New("dpiStmt_execute: ORA-04064: not executed, invalidated %s")
	ErrByName["TNSAdapterErr"] = errors.New("dpiStmt_execute: ORA-12560: TNS:protocol adapter error")
	ErrByName["snapshotTooOld"] = errors.New("dpiStmt_execute: ORA-01555: snapshot too old: rollback segment number %s with name \"%s\" too small")
	ErrByName["invalidCredentials"] = errors.New("dpiStmt_execute: ORA-01017: invalid username/password; logon denied")
	ErrByName["missingExpression"] = errors.New("dpiStmt_execute: ORA-00936: missing expression")
	ErrByName["invalidNumber"] = errors.New("dpiStmt_execute: ORA-01722: invalid number")
	ErrByName["PL/SQLValueErr"] = errors.New("dpiStmt_execute: ORA-06502: PL/SQL: numeric or value error%s")
	ErrByName["invalidCharacter"] = errors.New("dpiStmt_execute: ORA-00911: invalid character")
	ErrByName["invalidSQL"] = errors.New("dpiStmt_execute: ORA-00933: SQL command not properly ended")
	ErrByName["bug"] = errors.New("dpiStmt_execute: ORA-00600: internal error code, arguments: [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s]")
	ErrByName["invalidNumber"] = errors.New("dpiStmt_execute: ORA-01722: invalid number")
	ErrByName["EOFOnChannel"] = errors.New("dpiStmt_execute: ORA-03113: end-of-file on communication channel")
	ErrByName["maxCursors"] = errors.New("dpiStmt_execute: ORA-01000: maximum open cursors exceeded")
	ErrByName["coreDump"] = errors.New("dpiStmt_execute: ORA-07445: exception encountered: core dump [%s] [%s] [%s] [%s] [%s] [%s]")

	ErrByNum[1555] = errors.New("dpiStmt_execute: ORA-1555: snapshot too old: rollback segment number %s with name %s too small")

}

func GetErrByName (errName string) (error, bool) {

	if val, ok := ErrByName[errName]; ok {
		return val,true
	}

	return nil,false

}

func GetErrByNum (errNum int) (error, bool) {

	if val, ok := ErrByNum[errNum]; ok {
		return val,true
	}

	return nil,false

}

func Test () {

	errNamesToChk := [3]string{
		"uniqueConstraint",
		"coreDump",
		"unknown",
	}

	errNumsToCheck := [4]int{
		942,
		12,
		1555,
		6512,
	}

	fmt.Println("all is go")

	for _, errName := range errNamesToChk {

		if errVal, ok := GetErrByName(errName); ok {
			fmt.Println("err: ", errVal)
		} else {
			fmt.Println("ErrByName: That error is unknown")
		}
	}

	for _, errNum := range errNumsToCheck {

		if errVal, ok := GetErrByNum(errNum); ok {
			fmt.Println("err: ", errVal)
		} else {
			fmt.Println("ErrByNum: That error is unknown")
		}
	}
}


