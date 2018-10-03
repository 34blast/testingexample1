package testingexample1

import (
	"flag"
//	"fmt"
	"os"
	"testing"
)

var db struct {
	Url string
}

func TestMain(m *testing.M) {
	// Pretend to open our DB connection
	// this is like Setup run before all tests
	db.Url = os.Getenv("DATABASE_URL")
	if db.Url == "" {
		db.Url = "localhost:5432"
	}

	flag.Parse()
	exitCode := m.Run()

	// Pretend to close our DB connection
	// this part is like tearDown in Java
	cleanUp()
	
	// Exit
	os.Exit(exitCode)
}

func TestDatabase(t *testing.T) {
//	fmt.Println("TestDatabase, db.Url = ", db.Url)
}

func TestDatabase2(t *testing.T) {
//	fmt.Println("TestDatabase2, db.Url = ", db.Url)
}

func TestDatabaseError(t *testing.T) {
	// Pretend to use the db
//	fmt.Println("TestDatabase3, db.Url = ", db.Url)
//	t.Error(fmt.Sprint("Invalid DB URL"))

}

func cleanUp() {
	db.Url = ""
}