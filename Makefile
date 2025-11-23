
all: build up

build:
	docker-compose -f compose.yaml build

up:
	docker-compose -f compose.yaml up -d

down:
	docker-compose -f compose.yaml down

clean:
	docker-compose -f compose.yaml down -v --rmi all