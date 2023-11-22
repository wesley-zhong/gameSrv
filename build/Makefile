GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_WORLD=world
BIN_GAME=game
BIN_GATE_WAY=gateway
BIN_CLIENT=client
all: clean  build
restart: stop run
build:
	set GOARCH=amd64
    go env -w GOARCH=amd64
    set GOOS=linux
    go env -w GOOS=linux
	$(GOBUILD) -o $(BIN_WORLD)/$(BIN_WORLD)  ../world/cmd/world.go
	cp  -r ../world/configs   ./$(BIN_WORLD)/
	cp  ../world/Dockerfile  ./$(BIN_WORLD)/
	$(GOBUILD) -o $(BIN_GAME)/$(BIN_GAME)  ../game/cmd/game.go
	cp -r ../game/configs ./$(BIN_GAME)
	cp ../game/Dockerfile  ./$(BIN_GAME)/
	$(GOBUILD) -o $(BIN_GATE_WAY)/$(BIN_GATE_WAY) ../gateway/cmd/gateway.go
	cp -r ../gateway/configs ./$(BIN_GATE_WAY)
	cp ../gateway/Dockerfile  ./$(BIN_GATE_WAY)/
	$(GOBUILD) -o $(BIN_CLIENT)/$(BIN_CLIENT) ../client/cmd/client.go
	cp -r ../client/configs ./$(BIN_CLIENT)

clean:
	rm -rf $(BIN_WORLD)
	rm -rf $(BIN_GAME)
	rm -rf $(BIN_GATE_WAY)
	rm -rf $(BIN_CLIENT)
	rm -rf ./logs/*
run:
	nohup ./$(BIN_WORLD)/$(BIN_WORLD) >>./logs/world.log 2>&1 &
	sleep 1
	nohup ./$(BIN_GAME)/$(BIN_GAME) >>./logs/game.log 2>&1 &
	sleep 1
	nohup ./$(BIN_GATE_WAY)/$(BIN_GATE_WAY) >>./logs/gateway.log 2>&1 &
	sleep 1
	@echo "======================== server started "
	ps -aux | grep gateway
	ps -aux | grep game
	ps -aux | grep world
ps:
	ps -aux | grep gateway
	ps -aux | grep game
	ps -aux | grep world



