start:
	make stop
	docker-compose up --remove-orphans

start-no-cache:
	make stop
	docker-compose build --no-cache
	docker-compose up --remove-orphans

stop:
	docker-compose down

reset-db:
	docker exec -i db mysql -u demo -pdemo db < db_dump.sql

run-tests:
	make docker-build-tests

docker-build-tests:
	@docker build \
		--tag tests_builder \
		-f tests.dockerfile .
