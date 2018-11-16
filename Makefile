#! /usr/bin/make

##### Convenience targets ######

REPO:=github.com/TinyKitten/ComingServer
GAE_PROJECT:=coming-222616

init: depend bootstrap
gen: clean generate

build:
	@go build -o comingserver

depend:
	@which glide || go get -v github.com/Masterminds/glide
	@glide install

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client  -d $(REPO)/design

deploy:
	gcloud app deploy --project $(GAE_PROJECT) .


##### Database ######

DBUSERNAME:=comingserver
DBNAME:=comingserver
ENV:=development

migrate/init:
	mysql -u comingserver -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

migrate/dry:
	sql-migrate up -dryrun -env=$(ENV)
