postgres:
	docker run --name auth-service-postgres12 -p 6002:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=auth-service -d postgres:12-alpine

createdb:
	docker exec -it auth-service-postgres12 createdb --username=root --owner=root auth-service-database

dropdb:
	docker exec -it auth-service-postgres12 dropdb auth-service-database

migrateup:
	migrate -path migration -database "postgresql://root:auth-service@localhost:6002/auth-service-database?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:auth-service@localhost:6001/auth-service-database?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/main.go	

proto:
	for service in user authentication; do \
		protoc --proto_path=internal/grpc/$$service/proto --go_out=internal/grpc/$$service/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/grpc/$$service/pb --go-grpc_opt=paths=source_relative \
		internal/grpc/$$service/proto/*.proto; \
	done

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server proto
