# Generate
This generates all proto models and services. But logic should still be impelemnted by human.

Look at [Requrements](#requirements) and [Config](#config) for setup

Usage:
```bash
$ ./generate.sh source target-directory
```
give persmission if needed by doing `chmod a+x generate.sh`

* `source`: a proto source file
* `target-directory`: path to destination directory where a proto and proto-grpc file will be generated

# Requirements
## Linux package:
* protoc-gen-go: `sudo apt-get install protoc-gen-go`

## Go packages:
* proto-gen-go: `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
* proto-gen-go-grpc: `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

# Config
For protoc to correctly find needed go libs and use them, we also need a go bin directory in our path.
It's located at `/home/go/bin` usually. We can do it by doing `expor PATH=$PATH:$GOPATH/bin` or adding it to our .[zsh/bash].rc file