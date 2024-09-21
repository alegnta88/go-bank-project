createdb:
	docker exec -it postgres-db createdb --username=admin --owner=admin go_test

dropdb:
	docker exec -it postgres-db dropdb go_test

migrateup:
	migrate -path db/migration -database "postgresql://admin:secret@147.182.241.168:5432/go_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:secret@147.182.241.168:5432/go_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate