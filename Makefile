# Makefile for TheBlock
#
# This makefile contains commands for running docker, and frontend and backend servers. 
# We are running local development for database migrations on docker
#
#
#

postgres:
    docker run --name devdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it devdb createdb --username=root --owner=root the_block_local

dropdb:
    docker exec -it devdb dropdb the_block_local

migrateup:
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/the_block_local?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/the_block_local?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
