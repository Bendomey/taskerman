package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/handler"
	"github.com/Bendomey/task-assignment/graph"
	"github.com/Bendomey/task-assignment/repository"
	"github.com/Bendomey/task-assignment/utils"
)

const defaultPort = "8080"

func main() {
	// load environmwnt variables in here
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository, err := repository.NewPostgresqlRepository(os.Getenv("DATABASE_URL_LIVE"))
	if err != nil {
		log.Fatalln(err)
	}

	//incase of any errors close connection
	defer repository.Close()

	//create super admin if it doesn't exists
	saveAdminErr := utils.SaveAdminInfo(repository)
	if saveAdminErr != nil {
		log.Fatalf("Was unable to save admin. Logs here: %s", saveAdminErr)
	}

	//set port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv, err := graph.NewGraphqlServer(repository)
	if err != nil {
		log.Fatalf("An error occured %s", err)
	}

	log.Printf("Server active. Goto http://localhost:%s/graphql for GraphQL playground", port)
	http.Handle("/", handler.GraphQL(srv.ToExecutableSchema()))
	http.Handle("/graphql", handler.Playground("Taskerman", "/"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
