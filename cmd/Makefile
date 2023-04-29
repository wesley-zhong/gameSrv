GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_WORLD=world
BIN_GAME=game
BIN_GATE_WAY=gateway
all: clean  build
restart: stop run
build:
	$(GOBUILD) -o $(BIN_WORLD)  ../world/cmd/world.go
	$(GOBUILD) -o $(BIN_GAME) ../game/cmd/game.go
	$(GOBUILD) -o $(BIN_GATE_WAY) ../gateway/cmd/gateway.go
clean:
	rm -f $(BIN_WORLD)
	rm -f $(BIN_GAME)
	rm -f $(BIN_GATE_WAY)
run:
	nohup ./$(BIN_WORLD) &
	sleep 5
	nohup ./$(BIN_GAME) &
	sleep 5
	nohup ./$(BIN_GATE_WAY)&
stop:
	 ps -ef|grep gateway|grep -v grep|awk '{print $2}' | xargs kill -9
	 ps -ef|grep game|grep -v grep|awk '{print $2}' | xargs kill -9
	 ps -ef|grep world|grep -v grep|awk '{print $2}' | xargs kill -9
