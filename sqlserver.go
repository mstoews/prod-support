package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Requests struct {
	SR_ID int64
	Summary string

}

type CountSummary struct {
	Count int64
	Service_Request_Stat string
	Requested string 
}

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
	// Report(conn)
	RequestSummary(conn)	
	fmt.Printf("bye\n")
}

func RequestSummary (conn *sql.DB){
	
	selectDB, err := conn.Query("select count(*) as Count, Service_Request_Stat, Requested from ServiceRequests group by Service_Request_Stat, Requested, Requested")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer selectDB.Close()
	requests := CountSummary{}
	req := []CountSummary{}
	
	var Count int64
	var Service_Request_Stat string
	var Requested string

	for selectDB.Next(){
		err = selectDB.Scan(&Count, &Service_Request_Stat, &Requested)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		requests.Count = Count
		requests.Requested = Requested
		requests.Service_Request_Stat = Service_Request_Stat
		req = append(req,requests)
	}
	// fmt.Printf("%d", req[0].Count)
	fmt.Printf("%s %s %d\n", Requested, Service_Request_Stat, Count)
}


func Report(conn *sql.DB){
	
	selectDB, err := conn.Query("select SR_ID, Summary from ServiceRequests")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer selectDB.Close()
	requests := Requests{}
	req := []Requests{}
	var SR_ID int64
	var Summary string
	fmt.Printf("SR_ID\t Summary \n")
	for selectDB.Next(){
		err = selectDB.Scan(&SR_ID, &Summary)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		requests.SR_ID = SR_ID
		requests.Summary = Summary
		req = append(req,requests)
		fmt.Printf("%d\t%s\n", SR_ID, Summary)
	}
	
}

// connString := fmt.Sprintf("server=%s;port=%d;database=%s;Integrated Security=SSPI" , *server, *port, *database)
// docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Pass@word' -p 5433:1433 -d mcr.microsoft.com/mssql/server:2017-latest
	