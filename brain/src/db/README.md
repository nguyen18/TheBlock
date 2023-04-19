# Database Migrations
Before running migrations make sure you have golang-migrate installed:

#### MAC OS
```
brew install golang-migrate
```

## Please generate migrations files using the following commands:
### In this folder
```
migrate create -ext sql -dir ./migrations -seq $(REPLACE WITH MIGRATION FILE NAME)
```

### In /migrations
```
migrate create -ext sql -seq $(REPLACE WITH MIGRATION FILE NAME)
```

example: migrate create -ext sql -dir ./migrations -seq create_users_table

## Running migrations on Docker Postgres Database
This is already setup for you. Please back our to the root of theblock and run `make migrate-up` to run migrations directly on postgres docker database container. Make sure that the databsae is setup by first running `make create-docker`. Or run `make start` to start entire local database and local servers (migrations will automatically run on this as well).
