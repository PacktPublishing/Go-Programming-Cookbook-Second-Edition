package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter12/graphql/cards"
	"github.com/graphql-go/graphql"
)

func main() {
	// grab our schema
	schema, err := cards.Setup()
	if err != nil {
		panic(err)
	}

	// Query
	query := `
		{
			cards(value: "A"){
				value
				suit
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", rJSON)
}
