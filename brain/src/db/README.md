brew install golang-migrate

#in /migration
migrate create -ext sql -dir db/migration -seq init_schema

