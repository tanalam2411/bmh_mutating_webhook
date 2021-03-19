#!/bin/bash

hostIP=$(hostname -I | awk '{print $1}')

if [ "$1" == "up" ]; then
  echo "up..."
  echo $hostIP
  MY_IP=$hostIP docker-compose -f ./hack/kafka-docker-compose.yml up -d
  sleep 2
  docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic topic-1 --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
fi


if [ "$1" == "down" ]; then
  echo "down ..."
  MY_IP=$hostIP docker-compose -f ./hack/kafka-docker-compose.yml down
fi
