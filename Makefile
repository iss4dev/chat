postgresinit:
	docker run --name=chat -e POSTGRES_PASSWORD='psql' -p 5432:5432 -d --rm postgres

postgres:
	docker exec -it chat psql -U postgres

createdb:
	migrate -path ./schema -database 'postgres://postgres:psql@localhost:5432/postgres?sslmode=disable' up

dropdb:
	migrate -path ./schema -database 'postgres://postgres:psql@localhost:5432/postgres?sslmode=disable' down

.PHONY: postgresinit createdb dropdb
