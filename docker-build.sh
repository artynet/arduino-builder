#!/bin/bash -x

docker run --name abuilder --rm -v $(pwd):/root/arduino-builder \
	-w /root/arduino-builder golang:latest \
	bash -c "apt-get update && apt-get install -y tar bzip2 git zip && chmod +x *.sh && ./setup.sh && ./deploy.sh"
