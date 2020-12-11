serve:
	reflex -r '\.go' -s -- sh -c "go run server.go"

migrate:
	migrate -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable -path db/migrations up

migrate-fix:
	migrate -path db/migrations -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable force $(version)

migrate-rollover:
	migrate -database postgres://domey:akankobateng1@localhost:5432/taskerman?sslmode=disable -path db/migrations down