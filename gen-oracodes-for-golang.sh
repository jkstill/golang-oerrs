#!/usr/bin/env bash

# the format at this time returned for oracle error messages

#grep -E '^[[:digit:]]{5}' $ORACLE_HOME/rdbms/mesg/oraus.msg| tr -d '"' | awk -F', ' '{ print "dpiStmt_execute: ORA-"$1": "$3 }' 

declare mode=TEST
mode=BUILD

echo
if [[ $mode == 'TEST' ]]; then
	echo package main
else
	echo package oerrs
fi

echo

printf "%s\n" \
'import (
	"errors"
	"fmt" 
)'


# common errors

cat <<-EOF

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

EOF

echo "var ErrByName map[string]error"
echo "var ErrByNum map[int]error"
echo 
echo "func init() {"
echo 
echo "	ErrByName = make(map[string]error,100)"
echo "	ErrByNum = make(map[int]error,22000)"
# echo "	ErrByName["uniqueConstraint"] =errors.New("dpiStmt_execute: ORA-00001: unique constraint (%s.%s) violated")"
echo 

while read name msg
do
	echo "	ErrByName[$name] = errors.New($msg)"
done < <(
cat <<-EOF
	"uniqueConstraint"		"dpiStmt_execute: ORA-00001: unique constraint (%s.%s) violated"
	"invalidIdentifier"	"dpiStmt_execute: ORA-00904: %s: invalid identifier"
	"tableNotFound"			"dpiStmt_execute: ORA-00942: table or view does not exist"
	"badConnectID"			"dpiStmt_execute: ORA-12154: TNS:could not resolve the connect identifier specified"
	"badConnectPath"		"dpiStmt_execute: ORA-12198: TNS:could not find path to destination"
	"cannotConnect"			"dpiStmt_execute: ORA-12203: TNS:unable to connect to destination"
	"listenerFailed"		"dpiStmt_execute: ORA-12500: TNS:listener failed to start a dedicated server process"
	"badAddress"				"dpiStmt_execute: ORA-12533: TNS:illegal ADDRESS parameters"
	"noSuchDB"				"dpiStmt_execute: ORA-12545: Connect failed because target host or object does not exist"
	"lostContact"			"dpiStmt_execute: ORA-12547: TNS:lost contact"
	"connectionClosed"		"dpiStmt_execute: ORA-12537: TNS:connection closed"
	"storedProcErr"			"dpiStmt_execute: ORA-04063: %s has errors"
	"storedProcInvalid"	"dpiStmt_execute: ORA-04064: not executed, invalidated %s"
	"TNSAdapterErr"			"dpiStmt_execute: ORA-12560: TNS:protocol adapter error"
	"snapshotTooOld"		"dpiStmt_execute: ORA-01555: snapshot too old: rollback segment number %s with name \\\"%s\\\" too small"
	"invalidCredentials"	"dpiStmt_execute: ORA-01017: invalid username/password; logon denied"
	"missingExpression"	"dpiStmt_execute: ORA-00936: missing expression"
	"invalidNumber"			"dpiStmt_execute: ORA-01722: invalid number"
	"PL/SQLValueErr"		"dpiStmt_execute: ORA-06502: PL/SQL: numeric or value error%s"
	"invalidCharacter"		"dpiStmt_execute: ORA-00911: invalid character"
	"invalidSQL"				"dpiStmt_execute: ORA-00933: SQL command not properly ended"
	"bug"						"dpiStmt_execute: ORA-00600: internal error code, arguments: [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s], [%s]"
	"invalidNumber"			"dpiStmt_execute: ORA-01722: invalid number"
	"EOFOnChannel"			"dpiStmt_execute: ORA-03113: end-of-file on communication channel"
	"maxCursors"				"dpiStmt_execute: ORA-01000: maximum open cursors exceeded"
	"coreDump"				"dpiStmt_execute: ORA-07445: exception encountered: core dump [%s] [%s] [%s] [%s] [%s] [%s]"
EOF
)

# all available errors in a map[] by number

echo

#:<< 'COMMENT'

while read oranum oracode oramsg
do

:<<'COMMENT'
cat <<-EOF

############################################3
oranum: $oranum
oracode: $oracode
oramsg: $oramsg

EOF
COMMENT

	
	#oramsg=$(echo "$oramsg" | sed -e 's/\\/\\\\/g')
	oramsg=${oramsg//\\/\\\\}
	echo "	ErrByNum[$oranum] = errors.New(\"dpiStmt_execute: ORA-$oranum: $oramsg\")"
done < <(
	#grep -E '^01555' $ORACLE_HOME/rdbms/mesg/oraus.msg | 
	grep -E '^[[:digit:]]{5}' $ORACLE_HOME/rdbms/mesg/oraus.msg | 
	tr -d '[",]' | 
	awk  '{ 
		printf("%s ", $1+0); 
		printf("%s ",$1); 
		for (i=3;i<=NF;i++) {printf("%s ", $i)};
			printf("\n")}' 
)

#COMMENT

echo

echo "}"


printf "\n%s\n" \
'func GetErrByName (errName string) (error, bool) {

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
'

if [[ $mode == 'TEST' ]]; then
	echo 'func main () {'
else
	echo 'func Test () {'
fi

printf "\n%s\n" \
'	errNamesToChk := [3]string{
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

'

