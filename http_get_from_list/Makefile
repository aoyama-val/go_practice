SRC = $(wildcard *.go)
BIN = bin/main

$(BIN): $(SRC)
	go build -o $@ $(SRC)

clean:
	rm -rf $(BIN)
