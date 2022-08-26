test:
	@./go.test.sh
.PHONY: test

coverage:
	@./go.coverage.sh
.PHONY: coverage

build:
	CGO_ENABLED=0 go build -o bin/mvth ./cmd/mvthbot/app
.PHONY: build

docker-build:
	docker build -t migrate -f ./db/Dockerfile.multistage .
	docker build -t mvthbot -f ./deployments/Dockerfile.multistage .
.PHONY: docker-build

build-migrate:
	 CGO_ENABLED=0 go build -o bin/migrate  ./cmd/mvthbot/migrations/
.PHONY: build-migrate

clear:
	rm bin/*
	rm -rf pgdata
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
	migrate -path db/migrations -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose up 1
.PHONY: migrateup

migratedown:
	migrate -path db/migrations -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose down 1
.PHONY: migratedown

migrateset:
	migrate -path db/migrations -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose up
.PHONY: migrateset

migratedrop:
	migrate -path db/migrations -database "postgresql://adm@localhost:5432/mvthdb?sslmode=disable" -verbose down
.PHONY: migratedrop
