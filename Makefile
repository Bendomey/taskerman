serve:
	reflex -r '\.go' -s -- sh -c "go run server.go"

migrate:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up