package main

import (
	"github.com/jet/pkg"
	"github.com/jet/pkg/database"
	"log"
	"os"
)

// @zzmlearning
// some tools:
// https://golangci-lint.run/usage/quick-start/
func main() {
	// Q: why fs has to initialize with &database.OsFileStorage{} rather than database.OsFileStorage{}?
	// A: Short summary:
	//    An assignment to a variable of interface type is valid if the value being assigned implements the interface it is assigned to.
	//    It implements it if its method set is a superset of the interface.
	//    The method set of pointer types includes methods with both pointer and non-pointer receiver.
	//    The method set of non-pointer types only includes methods with non-pointer receiver.
	// todo:
	// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	fs := &database.OsFileStorage{}
	app := app.NewCliApp(fs)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
