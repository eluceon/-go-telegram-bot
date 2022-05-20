.DEFAULT_GOAL := docker-up

#######################################
##    protocol buffer compilation    ##
#######################################

PROTOC = protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative

.PHONY: protoc
protoc:
	$(PROTOC) ./api/quizpb/quiz.proto
	$(PROTOC) ./api/userpb/user.proto

.PHONY: clean-proto
clean-proto:
	rm -rf ./pkg/api/quizpb/*.go
	rm -rf ./pkg/api/userpb/*.go

.PHONY:docker-up
docker-up:
	docker-compose -f ./docker-compose.yml up

.PHONY:docker-down
docker-down:
	docker-compose -f ./docker-compose.yml down

.PHONY:docker-stop
docker-stop:
	docker-compose -f ./docker-compose.yml stop

.PHONY:docker-prune
docker-prune:
	docker system prune -af --volumes

######################
##    migrations    ##
######################

GOOSE = goose -dir migrations postgres "host=localhost port=54320 user=quizuser password=quizpass dbname=quizbot sslmode=disable"

.PHONY: migrate-create
migrate-create:
	goose -dir migrations create init sql

.PHONY: migrate-status
migrate-status:
	$(GOOSE) status

.PHONY: migrate-up
migrate-up:
	$(GOOSE) up

.PHONY: migrate-down
migrate-down:
	$(GOOSE) down

.PHONY: migrate-version
migrate-version:
	$(GOOSE) version

