package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/handler"
	"github.com/Bendomey/task-assignment/graph"
	"github.com/Bendomey/task-assignment/repository"
)

const defaultPort = "8080"

func main() {
	// load environmwnt variables in here
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//get repository in here
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		log.Fatalln("Please add a database url to your environment variables under the key: DATABASE_URL")
		os.Exit(1)
	}

	repository, err := repository.NewPostgresqlRepository(databaseurl)
	if err != nil {
		log.Fatalln(err)
	}

	//incase of any errors close connection
	defer repository.Close()

	//set port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv, err := graph.NewGraphqlServer(repository)
	if err != nil {
		log.Fatalf("An error occured %s", err)
	}

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	http.Handle("/", handler.GraphQL(srv.ToExecutableSchema()))
	http.Handle("/graphql", handler.Playground("Taskerman", "/"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
