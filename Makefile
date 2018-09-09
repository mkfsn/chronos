NAME = mkfsn/chronos

all:

build:
	docker build . -rm -t $(NAME)

push: login
	docker push $(NAME)

login:
	docker login -u=$(DOCKER_USERNAME) -p=$(DOCKER_PASSWORD)
