docker-build:
	docker build . --rm -t on2411/bmo-mutatingwebhook

docker-push:
	docker push on2411/bmo-mutatingwebhook