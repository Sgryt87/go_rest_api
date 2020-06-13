CREATE MIGRATION
migrate create -ext sql -dir migrations create_users

MIGRATE UP
migrate -path migrations -database "postgres://localhost/restapigo_dev?sslmode=disable" up