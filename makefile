RUN_PATH=cmd/web/main.go

run:
	go run ./$(RUN_PATH)

see:
	docker ps

delete:
	docker container prune

postgres:
	docker exec -it todoapp-postgres psql -U todoapp
