DEF_MONGO_URI=mongodb://localhost:27217
DEF_MONGO_DB=testdb
DEF_MONGO_EVENTS_COL=Events

AGENT_PORT=8080
WEB_PORT=8081

build_seeder:
	echo "Compiling Seeder"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/seeder-linux64 src/cmd/seeder/main.go

build_agent:
	echo "Compiling Deletion Agent"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/agent-linux64 src/cmd/agent/main.go

build_web:
	echo "Compiling Test Web API"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/web-linux64 src/cmd/web/main.go

build: build_seeder build_agent build_web

composition_start:
	(cd docker-compose && sh ./init-cluster.sh)

run_seeder: build_seeder
	chmod +x bin/seeder-linux64 && ./bin/seeder-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --accounts=10 --users=1000 --events=1000

run_agent: build_agent
	chmod +x bin/agent-linux64 && ./bin/agent-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --port=$(AGENT_PORT)

run_web: build_web
	chmod +x bin/web-linux64 && ./bin/web-linux64 --uri=$(DEF_MONGO_URI) --db=$(DEF_MONGO_DB) --col=$(DEF_MONGO_EVENTS_COL) --port=$(WEB_PORT)

docker_agent_build: build_agent
	docker build -t mongo-delete-agent-service -f Dockerfile.agent .

docker_agent_run: docker_agent_build
	docker run --name mongo-delete-agent-service -p 8080:80 -it mongo-delete-agent-service --uri=mongodb://host.docker.internal:27217 --db=testdb --col=Events

clean:
	rm bin/seeder-linux64