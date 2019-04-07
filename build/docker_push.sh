#!/bin/bash

VERSION_REGEX="/\d(.\d)+/"

IMAGE_TAG="latest"

if [[ $TRAVIS_TAG =~ $VERSION_REGEX ]]; then
    IMAGE_TAG=$TRAVIS_TAG
fi;

docker build -f docker/web-app/Dockerfile -t $WEB_TAG:$IMAGE_TAG

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin;

docker tag $WEB_TAG:$IMAGE_TAG $DOCKER_LOGIN/$WEB_TAG:$IMAGE_TAG;

docker push $DOCKER_LOGIN/$WEB_TAG:$IMAGE_TAG;