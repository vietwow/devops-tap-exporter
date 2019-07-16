#!/bin/bash

docker login && docker build -t vietwow/tap_exporter . && docker push vietwow/tap_exporter
