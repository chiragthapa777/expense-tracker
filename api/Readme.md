# for migration we are using golang-migrate
# command for the migration
# migration should be created in the migrations folder in following format
    - 0011_migration_name.up.sql
    - 0011_migration_name.down.sql
> migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/expense-tracker?sslmode=disable" up
> migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/expense-tracker?sslmode=disable" down