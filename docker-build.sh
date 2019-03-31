#!/bin/bash -x

DUSER=$(whoami)
DUID=${UID}

docker run --name abuilder --rm -v $(pwd):/root/arduino-builder \
	-w /root/arduino-builder golang:latest \
	bash -c "useradd ${DUSER} -u ${DUID} && \
	apt-get update && apt-get install -y tar bzip2 git zip && \
	chmod +x *.sh && ./setup.sh && ./deploy.sh && \
	chown ${DUSER}:${DUSER} -R distrib/ && \
	chown ${DUSER}:${DUSER} -R arduino-builder/"
