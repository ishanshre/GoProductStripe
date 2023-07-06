#!make

include .env

DB_URL=postgresql://${m_db_username}:${m_db_password}@localhost:5432/${m_db_dbname}?sslmode=disable

createDBContainer:
	docker run --name goProductDB -e POSTGRES_USER=${m_db_username} -e POSTGRES_PASSWORD=${m_db_password} -p 5432:5432 -d postgres

createDBPGadmin4Container:
	docker run --name goProductPgAdmin -p 5050:80 -e 'PGADMIN_DEFAULT_EMAIL=admin@admin.com' -e 'PGADMIN_DEFAULT_PASSWORD=admin' -d dpage/pgadmin4

createRedisContainer:
	docker run -d --name goProductRedis -p 6379:6379 redis:latest 

startContainer:
	docker start goProductDB goProductPgAdmin goProductRedis

stopContainer:
	docker stop goProductDB goProductPgAdmin goProductRedis

buildAPI:
	go build -o api cmd/api/* && ./api
buildWEB:
	go build -o web cmd/web/*.go && ./web
