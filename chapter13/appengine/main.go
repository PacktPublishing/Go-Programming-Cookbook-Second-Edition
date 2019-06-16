package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	log.SetOutput(os.Stderr)

	// Set this in app.yaml when running in production.
	projectID := os.Getenv("GCLOUD_DATASET_ID")

	datastoreClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	c := Controller{datastoreClient}

	http.HandleFunc("/", c.handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
