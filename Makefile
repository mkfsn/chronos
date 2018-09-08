NAME = mkfsn/chronos


all:

build:
	docker build . -t $(NAME)

push:
	docker push $(NAME)
