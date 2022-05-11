package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"	//Library required to establish Database
)

const (
	host	= "localhost"
	port	= 5432
	user	= "user"	
	pass	= "pass"
	schema	= "schema"
)

var (
	Client sql.DB
)
	

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s schema=%s sslmode=disable",
		host, port, user, pass, schema)
	
	Client, err := sql.Open("postgres", connectionString)
	if err != nil {
		//Exception
		fmt.Println("Failed to open database connection")
	}

	err = Client.Ping()
	if err != nil {
		//Exception
		fmt.Println("Failed to reach database")
	}

	fmt.Println("Successfully connected!")
}