build_seeder:
	echo "Compiling Seeder"
	GOARCH=amd64 GOOS=linux go build -o bin/seeder-linux64 cmd/seeder/main.go

composition_start:
	(cd docker-compose && sh ./init-cluster.sh)

clean:
	rm bin/seeder-linux64