#!/bin/bash
printf "Init cluster step 1: Start all of the containers\n"
docker-compose up -d --force-recreate

printf "\nWaiting 20 seconds for all containers to start\n"
sleep 20s

printf "\nInit cluster step 2.1: Initialize config servers rs\n"
docker-compose exec -T configsvr01 sh -c "mongo < /scripts/init-configserver.js"

printf "\nInit cluster step 2.2: Initialize shard 1 rs\n"
docker-compose exec -T shard01-a sh -c "mongo < /scripts/init-shard01.js"

printf "\nInit cluster step 2.3: Initialize shard 2 rs\n"
docker-compose exec -T shard02-a sh -c "mongo < /scripts/init-shard02.js"

printf "\nInit cluster step 2.4: Initialize shard 3 rs\n"
docker-compose exec -T shard03-a sh -c "mongo < /scripts/init-shard03.js"

printf "\nWaiting 50 seconds for all replica sets to elect their primaries\n"
sleep 50s

printf "\nInit cluster step 3: Initializing the router\n"
docker-compose exec -T router01 sh -c "mongo < /scripts/init-router.js"
printf "\n"
