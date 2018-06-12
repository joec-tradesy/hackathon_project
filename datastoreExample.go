package main

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"github.com/joec-tradesy/hackathon_project/immutableDatastore"
	"golang.org/x/net/context"
	"log"
)

const ProjectID = "tradesy-sandbox"

// func f(ctx context.Context) {

// 	// ...
// 	employee := &Employee{
// 		FirstName: "Antonio",
// 		LastName:  "Salieri",
// 		HireDate:  time.Now(),
// 	}
// 	employee.AttendedHRTraining = true

// 	//Creates a client.
// 	client, err := datastore.NewClient(ctx, ProjectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// Sets the kind for the new entity.
// 	kind := "HackathonEmployee"
// 	// Sets the name/ID for the new entity.
// 	name := "sampleEmployee"
// 	// Creates a Key instance.
// 	employeeKey := datastore.NameKey(kind, name, nil)

// 	if _, err := client.Put(ctx, employeeKey, employee); err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	// ...
// }

type Employee struct {
	immutableDatastore.ImmutableEntity
	FirstName string
	LastName  string
}

func main() {
	ctx := context.Background()

	//Creates a client.
	client, err := datastore.NewClient(ctx, ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var e = Employee{}
	e.FirstName = "Mike"
	e.LastName = "Hemelberg"
	e.KindName = "HackathonEmployee"
	e.Client = *client

	fmt.Println("1", e)

	err = e.Create()
	if err != nil {
		fmt.Println("ERROR %v", err)
	}

	// err = e.Get("sampleEmployee")
	// if err != nil {
	// 	fmt.Println("ERROR %v", err)
	// }

	// fmt.Println(e.FirstName)

}
