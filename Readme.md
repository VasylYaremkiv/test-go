Start server:
go get github.com/lib/pq
go get github.com/gorilla/mux
go get -v github.com/rubenv/sql-migrate/...

export GOPATH="/home/vasyly/work/test-go:$GOPATH"
bin/sql-migrate up
go install app
./bin/setup

Setup
1. Create database
2. 

