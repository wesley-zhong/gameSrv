GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_ROUTER=router
BIN_GAME=game
BIN_GATE_WAY=gateway
BIN_LOGIN=login
BIN_CLIENT=client
all: clean  build
restart: stop run
docker:  clean build docker_build
build:
	set GOARCH=amd64
	go env -w GOARCH=amd64
	set GOOS=linux
	go env -w GOOS=linux
	$(GOBUILD) -o $(BIN_ROUTER)/$(BIN_ROUTER)  ../router/cmd/router.go
	cp  -r ../router/configs   ./$(BIN_ROUTER)/
	cp  ../router/Dockerfile  ./$(BIN_ROUTER)/
	$(GOBUILD) -o $(BIN_GAME)/$(BIN_GAME)  ../game/cmd/game.go
	cp -r ../game/configs ./$(BIN_GAME)
	cp ../game/Dockerfile  ./$(BIN_GAME)/
	$(GOBUILD) -o $(BIN_GATE_WAY)/$(BIN_GATE_WAY) ../gateway/cmd/gateway.go
	cp -r ../gateway/configs ./$(BIN_GATE_WAY)
	cp ../gateway/Dockerfile  ./$(BIN_GATE_WAY)/
	$(GOBUILD) -o $(BIN_LOGIN)/$(BIN_LOGIN) ../login/cmd/login.go
	cp -r ../login/configs ./$(BIN_LOGIN)
	cp ../login/Dockerfile  ./$(BIN_LOGIN)/

	$(GOBUILD) -o $(BIN_CLIENT)/$(BIN_CLIENT) ../client/cmd/client.go
	cp -r ../client/configs ./$(BIN_CLIENT)
clean:
	rm -rf $(BIN_ROUTER)
	rm -rf $(BIN_GAME)
	rm -rf $(BIN_GATE_WAY)
	rm -rf $(BIN_CLIENT)
	rm -rf ./logs/*
run:
	nohup ./$(BIN_ROUTER)/$(BIN_ROUTER) >>./logs/router.log 2>&1 &
	sleep 1
	nohup ./$(BIN_GAME)/$(BIN_GAME) >>./logs/game.log 2>&1 &
	sleep 1
	nohup ./$(BIN_GATE_WAY)/$(BIN_GATE_WAY) >>./logs/gateway.log 2>&1 &
	sleep 1
	@echo "======================== server started "
	ps -aux | grep gateway
	ps -aux | grep game
	ps -aux | grep router
stop:
	ps -ef|grep gateway|grep -v grep|awk '{print $2}' | xargs kill -9
	ps -ef|grep game|grep -v grep|awk '{print $2}' | xargs kill -9
	ps -ef|grep router|grep -v grep|awk '{print $2}' | xargs kill -9
	ps -aux | grep game
	ps -aux | grep gateway
	ps -aux | grep router
ps:
	ps -aux | grep gateway
	ps -aux | grep game
	ps -aux | grep router
docker_build:
	cd gateway; docker build --tag docker-gateway .
	cd game; docker build --tag docker-game .
	cd router; docker build --tag docker-router .
	cd login; docker build --tag docker-login .
docker_run:
	docker run -d -p 9101:9101 --name gateway  docker-gateway
	docker run -d -p 9102:9102 --name game  docker-game
	docker run -d -p 9103:9103 --name router  docker-router



