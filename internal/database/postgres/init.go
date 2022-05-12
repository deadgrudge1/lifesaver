package postgres

import (
	"log"
	"os"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
)

var (
	host	= os.Getenv("POSTGRES_DB_HOST")
	port	= os.Getenv("POSTGRES_DB_PORT")
	user	= os.Getenv("POSTGRES_DB_USER")	
	pass	= os.Getenv("POSTGRES_DB_PASS")
	schema	= os.Getenv("POSTGRES_DB_NAME")
)

var (
	Client sql.DB
)

func flyway() error {
	connectionString := fmt.Sprintf("postgres://%s:%d/%s?sslmode=enable",
		host, port, schema, pass, schema)

	//Initiate Connection
	db, err := sql.Open("postgres", connectionString)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if(err != nil) {
		log.Println("[FLYWAY] Failed to create instance for Postgres Connection")
		return err
	}

	//Migrate
    migration, err := migrate.NewWithDatabaseInstance("file:///db/migrations", "postgres", driver)
	if(err != nil) {
		log.Println("[FLYWAY] Failed to create Database Instance")
		return err;
	}

    migration.Up()

	//Close flyway connection
	db.Close()

	return nil
}
	

func init() {
	//Create tables from migration
	flyway()

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s schema=%s sslmode=disable",
		host, port, user, pass, schema)
	
	//Open Connection
	Client, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("[DATA_BASE] Failed to open database connection to Postgres")
	}

	//Ping Database
	err = Client.Ping()
	if err != nil {
		log.Println("[DATA_BASE] Failed to reach database - Postgres")
	} else {
		log.Println("[DATA_BASE] Successfully Connected - Postgres")
	}
}