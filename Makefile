build_seeder:
	echo "Compiling Seeder"
	GOARCH=amd64 GOOS=linux go build -o bin/seeder-linux64 cmd/seeder/main.go

build_agent:
	echo "Compiling Deletion Agent"
	GOARCH=amd64 GOOS=linux go build -o bin/agent-linux64 cmd/agent/main.go

build: build_seeder build_agent

composition_start:
	(cd docker-compose && sh ./init-cluster.sh)

clean:
	rm bin/seeder-linux64