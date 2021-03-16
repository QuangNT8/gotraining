all: build-estest

init-project:
	@go mod init zmq

install-pkg: #install tools
	@go get -u github.com/pebbe/zmq4
	@go get -u gopkg.in/yaml.v2
	
build-estest:
	@go build -o estest estest.go

clean:
	@go clean
	rm -rf estest
	# rm -rf ./build/client
run:
	./estest
