package main
  
import (
    "fmt"
	 "context"
	 //"time"
	 "log"
    "database/sql"
	 "github.com/jkstill/golang-oerrs"
    _ "github.com/godror/godror"
)

var pool *sql.DB // Database connection pool.
var ctx context.Context // not using this yet

func main(){

	//id := 1
	var err error

    pool, err = sql.Open("godror", "scott/tiger@myserver/pdb1")
    if err != nil {
        fmt.Println(err)
        return
    }

	 Ping()

    defer pool.Close()

    rows,err := pool.Query("select id,sysdate from mytable")

	if err != nil {

		//if err.Error() == "dpiStmt_execute: ORA-00942: table or view does not exist" {
		chkErrVal, _ :=  oerrs.GetErrByNum(942)
		//fmt.Println("      err:",err.Error())
		//fmt.Println("chkErrVal:",chkErrVal.Error())

		if err.Error() ==  chkErrVal.Error() {
			fmt.Println("That Table does not exist")
			createTable()
			fmt.Println("Table mytable created - please try again")
			return
		} else {
			fmt.Println("Error running query")
			fmt.Println(err)
			return
		}
	}

	defer rows.Close()

    var thedate string
	 var myID int
    for rows.Next() {

        rows.Scan(&myID,&thedate)
		  fmt.Printf("ID: %4d The date is: %s\n", myID, thedate)
    }

	 Query(2)
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping() {

	if err := pool.Ping(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return	
}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func Query(id int) {

	var name string
	err := pool.QueryRow("select m.id from mytable m where m.id = :id", sql.Named("id", id)).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("id=", name)

	return
}

func createTable() {

	sqltext := "create table mytable (id number, c1 varchar2(30))"

	pool.Exec(sqltext)
	pool.Exec("insert into mytable(id,c1) values(:id,'testing')", sql.Named("id",1))
	pool.Exec("insert into mytable(id,c1) values(:id,'testing')", sql.Named("id",2))

	return

}


