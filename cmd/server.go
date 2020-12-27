package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/handler"
	"github.com/Bendomey/task-assignment/graph"
	"github.com/Bendomey/task-assignment/repository"
	"github.com/Bendomey/task-assignment/utils"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

//defining the graphql handler
func graphqlHandler(srv *graph.Resolver) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(srv.ToExecutableSchema())

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("Taskerman", "/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// load environmwnt variables in here
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository, err := repository.NewPostgresqlRepository(os.Getenv("DATABASE_URL"))
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

	//setting up gin
	r := gin.Default()
	r.POST("/", graphqlHandler(srv))
	r.GET("/graphql", playgroundHandler())
	r.Run()
}
