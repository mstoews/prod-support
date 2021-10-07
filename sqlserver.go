package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "Pass@word ", "the database password")
	port     *int = flag.Int("port", 5433, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database      = flag.String("database", "prod_support", "database name")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
	}

	
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user=%s;password=%s" , *server, *port, *database, *user, *password)
	
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
	Report(conn)	
	fmt.Printf("bye\n")
}

func Report(conn *sql.DB){
	selectDB, err := conn.Query("select SR_ID, Summary from ServiceRequests")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer selectDB.Close()

	var SR_ID int64
	var Summary string
	fmt.Printf("SR_ID\t Summary \n")
	for selectDB.Next(){
		err = selectDB.Scan(&SR_ID, &Summary)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		fmt.Printf("%d\t%s\n", SR_ID, Summary)
	}
	
}

// connString := fmt.Sprintf("server=%s;port=%d;database=%s;Integrated Security=SSPI" , *server, *port, *database)
// docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Pass@word' -p 5433:1433 -d mcr.microsoft.com/mssql/server:2017-latest
	