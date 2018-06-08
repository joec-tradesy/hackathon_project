package main

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Employee struct {
	FirstName          string
	LastName           string
	HireDate           time.Time
	AttendedHRTraining bool
}

const ProjectID = "tradesy-sandbox"

func f(ctx context.Context) {

	// ...
	employee := &Employee{
		FirstName: "Antonio",
		LastName:  "Salieri",
		HireDate:  time.Now(),
	}
	employee.AttendedHRTraining = true

	//Creates a client.
	client, err := datastore.NewClient(ctx, ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "HackathonEmployee"
	// Sets the name/ID for the new entity.
	name := "sampleEmployee"
	// Creates a Key instance.
	employeeKey := datastore.NameKey(kind, name, nil)

	if _, err := client.Put(ctx, employeeKey, employee); err != nil {
		fmt.Println("ERROR")
	}
	// ...
}

func main() {
	ctx := context.Background()
	f(ctx)
}
