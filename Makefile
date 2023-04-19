# Makefile for TheBlock
#
# This makefile contains commands for running docker, and frontend and backend servers. 
# We are running local development for database migrations on docker
#
# make start to start entire local dev env
#

### VARIABLES
MIGRATIONS_PATH=brain/src/db/migrations
POSTGRES_PORT=5432

REACT_PORT=3000
REACT_PID=$(lsof -ti :$(REACT_PORT))

TESTDB_NAME=testdb
TESTDB_HOST=localhost
TESTDB_USER=root
TESTDB_PASS=secret

.PHONY: run-docker clean-docker create-docker drop-docker migrate-up migrate-down run-react stop-react start clean

start: run-docker run-react

clean: clean-docker stop-react

run-docker: create-docker migrate-up

clean-docker: migrate-down drop-docker

create-docker:
	docker-compose up -d

drop-docker:
	docker-compose down

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://$(TESTDB_USER):$(TESTDB_PASS)@$(TESTDB_HOST):$(POSTGRES_PORT)/$(TESTDB_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://$(TESTDB_USER):$(TESTDB_PASS)@$(TESTDB_HOST):$(POSTGRES_PORT)/$(TESTDB_NAME)?sslmode=disable" -verbose down

run-react:
	$(shell cd ./react_app/the-block && npm start)

stop-react:
	$(shell cd ./react_app/the-block && lsof -ti :$(REACT_PORT) | xargs kill)
