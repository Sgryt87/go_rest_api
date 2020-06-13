###CREATE MIGRATION
```migrate create -ext sql -dir migrations create_users```

####MIGRATE UP:
```
DEV:
 - migrate -path migrations -database "postgres://localhost/restapigo_dev?sslmode=disable" up
TEST_DB:
 - migrate -path migrations -database "postgres://localhost/restapigo_test?sslmode=disable" up
```