package postgres

import (
	"log"
	"os"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
    // "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var (
	host	= os.Getenv("POSTGRES_DB_HOST")
	port	= os.Getenv("POSTGRES_DB_PORT")
	user	= os.Getenv("POSTGRES_DB_USER")	
	pass	= os.Getenv("POSTGRES_DB_PASS")
	schema	= os.Getenv("POSTGRES_DB_NAME")
)

var (
	Client *sql.DB
	
	// go:embed lifesaver/db/migration/*.sql
	fs embed.FS
)



func flyway() error {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, schema)

	// //Initiate Connection
	// db, err := sql.Open("postgres", connectionString)
	// if(err != nil) {
	// 	log.Println("Failed to open connection : ", err)
	// 	return err
	// }

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if(err != nil) {
	// 	log.Println("[FLYWAY] Failed to create instance for Postgres Connection : ", err)
	// 	return err
	// }

	d, err := iofs.New(fs, "lifesaver/db/migration") // Get migrations from sql folder
	if err != nil {
		log.Fatal(err)
	}

	//Migrate
    // migration, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://lifesaver/db/migarion"), "postgres", driver)
	migration, err := migrate.NewWithSourceInstance("iofs", d, connectionString)
	if(err != nil) {
		log.Println("[FLYWAY] Failed to create Database Instance : ", err)
		return err;
	}

    migration.Up()

	//Close flyway connection
	// db.Close()

	return nil
}
	

func init() {
	//Create tables from migration
	// flyway()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, schema)
	
	//Open Connection
	var err error
	Client, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("[DATABASE] Failed to open database connection to Postgres : ", err)
	}

	//Ping Database
	err = Client.Ping()
	if err != nil {
		log.Println("[DATABASE] Failed to reach database - Postgres : ", err)
	} else {
		log.Println("[DATABASE] Successfully Connected - Postgres")
	}
}