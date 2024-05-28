postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todo_list

dropdb:
	docker exec -it postgres12 dropdb todo_list

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/todo_list?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/todo_list?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown