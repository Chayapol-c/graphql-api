package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	gql "graphql-tutorial/graphql"
)

func main() {

	rootQuery := gql.RootQuery // Assuming you have a root query defined in schema.go
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		panic(err)
	}
	// Create a GraphQL handler for HTTP requests
	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false, // Disable GraphiQL for subscriptions endpoint
		Playground: true,
	})

	// Serve GraphQL API at /graphql endpoint
	http.Handle("/graphql", graphqlHandler)

	// Start the HTTP server
	fmt.Printf(`Server is running at http://localhost:%v/graphql`, 4000)
	http.ListenAndServe(fmt.Sprintf(`:%v`, 4000), nil)

}
