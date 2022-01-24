DEF_MONGO_URI=mongodb://localhost:27217
DEF_MONGO_DB=testdb
DEF_MONGO_COL=testcol

build_seeder:
	echo "Compiling Seeder"
	GOARCH=amd64 GOOS=linux go build -o bin/seeder-linux64 cmd/seeder/main.go

build_agent:
	echo "Compiling Deletion Agent"
	GOARCH=amd64 GOOS=linux go build -o bin/agent-linux64 cmd/agent/main.go

build: build_seeder build_agent

composition_start:
	(cd docker-compose && sh ./init-cluster.sh)

run_seeder: build_seeder
	chmod +x bin/seeder-linux64 && ./bin/seeder-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_COL) --accounts=10 --users=1000000

run_agent: build_agent
	chmod +x bin/agent-linux64 && ./bin/agent-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_COL)

clean:
	rm bin/seeder-linux64