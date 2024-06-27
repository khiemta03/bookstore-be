postgres:
	docker run --name book-service-postgres12 -p 6003:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=book-service -d postgres:12-alpine

createdb:
	docker exec -it book-service-postgres12 createdb --username=root --owner=root book-service-database

dropdb:
	docker exec -it book-service-postgres12 dropdb book-service-database

migrateup:
	migrate -path migration -database "postgresql://root:book-service@localhost:6003/book-service-database?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:book-service@localhost:6003/book-service-database?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/main.go	

proto:
	protoc --proto_path=internal/grpc/proto/book:internal/grpc/proto/google/type \
       --go_out=internal/grpc/gen/book --go_opt=paths=source_relative \
       --go-grpc_out=internal/grpc/gen/book --go-grpc_opt=paths=source_relative \
       internal/grpc/proto/book/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
