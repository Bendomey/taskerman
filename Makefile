# for running application
run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server.go"

# for db migration
migrate:
	migrate -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable -path db/migrations up

migrate-fix:
	migrate -path db/migrations -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable force $(version)

migrate-rollover:
	migrate -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable -path db/migrations down

# for generating grapqhl
generate-graph:
	go run github.com/99designs/gqlgen generate

# for containerizing to docker
build-docker:
	docker build -t 0545526664/taskerman .