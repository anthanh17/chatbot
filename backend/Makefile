# Setup postgres database docker
postgres:
	docker run --name chatbotai -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d ankane/pgvector

createdb:
	docker exec -it chatbotai createdb --username=postgres --owner=postgres annt-vector

dropdb:
	docker exec -it chatbotai dropdb annt-vector

# Start server http
server:
	go run main.go

.PHONY: postgres createdb dropdb server
