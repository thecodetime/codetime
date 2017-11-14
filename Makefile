all: dev

dev:
	$(MAKE) build-watch & gin -p 4000

prod: build
	go install
	codetime -env production

build:
	gulp build

build-watch:
	gulp

reload:
	git pull origin master
	$(MAKE) build
	go install
	supervisorctl reload

