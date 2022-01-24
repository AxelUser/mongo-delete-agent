DEF_MONGO_URI=mongodb://localhost:27217
DEF_MONGO_DB=testdb
DEF_MONGO_EVENTS_COL=Events

AGENT_PORT=8080
WEB_PORT=8081

build_seeder:
	echo "Compiling Seeder"
	GOARCH=amd64 GOOS=linux go build -o bin/seeder-linux64 cmd/seeder/main.go

build_agent:
	echo "Compiling Deletion Agent"
	GOARCH=amd64 GOOS=linux go build -o bin/agent-linux64 cmd/agent/main.go

build_web:
	echo "Compiling Test Web API"
	GOARCH=amd64 GOOS=linux go build -o bin/web-linux64 cmd/web/main.go

build: build_seeder build_agent

composition_start:
	(cd docker-compose && sh ./init-cluster.sh)

run_seeder: build_seeder
	chmod +x bin/seeder-linux64 && ./bin/seeder-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --accounts=10 --users=1000 --events=1000

run_agent: build_agent
	chmod +x bin/agent-linux64 && ./bin/agent-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --port=$(AGENT_PORT)

run_web: build_web
	chmod +x bin/web-linux64 && ./bin/web-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --port=$(WEB_PORT)

clean:
	rm bin/seeder-linux64