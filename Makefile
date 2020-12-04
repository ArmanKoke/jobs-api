TARGET=jobs-api

swagger:
	swagger generate spec --compact -o ./swagger.json -m

# change container name, user, db
table:
	docker exec -i postgre psql -U postgres postgres < table.sql
psql:
	docker exec -ti postgre psql -U postgres dev

# following steps for building
build:
	go build main.go

docker:
	docker build -t $(TARGET) .

docker_run:
	docker rm $(TARGET) && docker run -p 8080:8080 --env-file .env --name $(TARGET) $(TARGET)

docker_id:
	docker ps -aqf "name=$(TARGET)"

container_ip:
	docker inspect --format '{{ .NetworkSettings.IPAddress }}' $(TARGET)