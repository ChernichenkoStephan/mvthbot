test:
	@./go.test.sh
.PHONY: test

coverage:
	@./go.coverage.sh
.PHONY: coverage

build:
	CGO_ENABLED=0 go build ./cmd/mvthbot
	mv mvthbot ./bin
.PHONY: build

clear:
	rm mvthbot
.PHONY: clear

check_generated: generate
	git diff --exit-code
.PHONY: check_generated

createdb:
	createdb --username=adm --owner=adm mvthdb
.PHONY: createdb

dropdb:
	dropdb -f mvthdb
.PHONY: dropdb

migrateup:
	migrate -path db/migration -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose up 1
.PHONY: migrateup

migratedown:
	migrate -path db/migration -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose down 1
.PHONY: migratedown

migrateset:
	migrate -path db/migration -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose up
.PHONY: migrateset

migratedrop:
	migrate -path db/migration -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose down
.PHONY: migratedrop
