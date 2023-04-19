# Makefile for TheBlock
#
# This makefile contains commands for running docker, and frontend and backend servers. 
# We are running local development for database migrations on docker
#
#
#

### VARIABLES
MIGRATIONS_PATH=brain/src/db/migrations
POSTGRES_PORT=5432
REACT_PORT=3000
REACT_PID=$(lsof -ti :$(REACT_PORT))
LOCAL_TESTDB_NAME=testdb

.PHONY: start-db clean-db create-db drop-db migrate-up migrate-down run-react stop-react

start-db: create-db migrate-up

clean-db: migrate-down drop-db

create-db:
	docker-compose up -d

drop-db:
	docker-compose down

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://root:secret@localhost:$(POSTGRES_PORT)/$(LOCAL_TESTDB_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://root:secret@localhost:$(POSTGRES_PORT)/$(LOCAL_TESTDB_NAME)?sslmode=disable" -verbose down

run-react:
	$(shell cd ./react_app/the-block && npm start)

stop-react:
	$(shell cd ./react_app/the-block && kill $(REACT_PID))
