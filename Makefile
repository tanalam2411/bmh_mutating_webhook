docker-build:
	docker build . --rm -t on2411/bmo-mutatingwebhook

docker-push:
	docker push on2411/bmo-mutatingwebhook

kafka-cluster-up:
	./hack/kafka-cluster.sh up

kafka-cluster-down:
	./hack/kafka-cluster.sh down

pkg-kafka-test:
	 ./hack/kafka-cluster.sh up \
 	&& sleep 2 \
 	&& echo "Running kafka tests..." \
 	&& go clean -testcache \
 	&& go test ./pkg/kafka \
 	&& sleep 10 \
 	&& ./hack/kafka-cluster.sh down
