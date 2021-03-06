# GO111MODULE := on

empty:
	go run cmd/modb/*.go

help:
	go run cmd/modb/*.go --help

start:
	go run cmd/modb/*.go start --store=mo.db

start-help:
	go run cmd/modb/*.go start --help

test:
	go test -v github.com/modb-io/modb/store

build:
	go build cmd/modb/modb.go

client:
	redis-cli -p 6380

dump:
	boltdb-dump mo.db

clean:
	rm -rf mo.db
