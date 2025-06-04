postgres:
	docker run --name postgres1 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:17.4-alpine3.21

createdb:
	docker exec -it postgres1 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres1 dropdb --username=root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@postgress:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:root@postgress:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:root@postgress:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@postgress:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	 mockgen -package mockdb -destination db/mock/store.go github.com/yuki/simplebank/db/sqlc/gen Store

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 sqlc test mock