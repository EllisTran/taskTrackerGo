.PHONY: all test clean

user     = ellistran 
binary   = taskgo 
version  = 0.0.1
build	   = $(shell git rev-parse HEAD)
ldflags  = -ldflags "-X 'github.com/$(user)/$(binary)/command.version=$(version)'
ldflags += -X 'github.com/$(user)/$(binary)/command.build=$(build)'"

all:
	go build -o $(binary) $(ldflags)

test:
	go test ./... -cover -coverprofile c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm -rf $(binary) c.out coverage.html

