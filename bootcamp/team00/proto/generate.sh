#!/bin/bash

SRC=$1
TAR=$2
function generate() {
	if [[ $SRC ]] && [[ $TAR ]]
	then
		echo "generating into $TAR for $SRC"

		protoc \
		--proto_path=. \
		--go_out=$TAR \
		--go-grpc_out=$TAR \
		$SRC
	else
		echo -e "Target or source is missing.\nUsage:\n./generate.sh file.proto ../some/destionation"
	fi
}

generate
