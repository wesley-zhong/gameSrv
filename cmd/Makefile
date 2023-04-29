GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_WORLD=world
BIN_GAME=game
BIN_GATE_WAY=gateway
all: clean  build
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
	nohup ./$(BIN_GAME) &
	nohup ./$(BIN_GATE_WAY)&
