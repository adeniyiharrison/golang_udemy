package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// used to extend database/sql available drivers
	_ "github.com/lib/pq"
)

var (
	axiomStageHost     = os.Getenv("LOCAL_USER_AXIOM_STAGE_HOST")
	axiomStageDBName   = os.Getenv("LOCAL_USER_AXIOM_STAGE_DBNAME")
	axiomStageUser     = os.Getenv("LOCAL_USER_AXIOM_STAGE_USER")
	axiomStagePassword = os.Getenv("LOCAL_USER_AXIOM_STAGE_PWD")
	axiomStagePort     = os.Getenv("LOCAL_USER_AXIOM_STAGE_PORT")
)

func checkErrors(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func returnConnectionString() string {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		axiomStageHost, axiomStagePort, axiomStageUser, axiomStagePassword, axiomStageDBName,
	)
	return connString
}

// ReturnDatabase returns Axiom DB object
func ReturnDatabase() *sql.DB {

	axiomStageConnString := returnConnectionString()
	db, err := sql.Open("postgres", axiomStageConnString)
	checkErrors(err)

	err = db.Ping()
	checkErrors(err)

	log.Println("successful connection to axiom db")

	return db

}
