all: build-mtserver

init-project:
	@go mod init zmq

install-pkg: #install tools
	@go get -u github.com/pebbe/zmq4
	@go get -u gopkg.in/yaml.v2
	
build-mtserver:
	@go build -o mtserver mtserver.go

clean:
	@go clean
	rm -rf mtserver
	# rm -rf ./build/client
run:
	./mtserver
