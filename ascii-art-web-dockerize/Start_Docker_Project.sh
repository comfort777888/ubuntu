#!/bin/bash
cd cmd
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
cd ..
echo "BUILDING DOCKER IMAGE WITH NAME <ascii-art-web>"
docker stop ascii-web
docker rm ascii-web
docker rmi ascii-art-web
docker build -t ascii-art-web .
echo "BUILDING AND RUNNING DOCKER CONTAINER WITH NAME <ascii-web> ON PORT 8080"
docker run -d -p 8080:8080 --name ascii-web ascii-art-web
echo "LISTS OF DOCKER IMAGES"
docker images 
echo "LISTS OF ALL EXISTING DOCKER CONTAINERS"
docker ps -a
echo "HELP TIPS: 1)TO PRUNES DOCKER'S IMAGES, CONTAINERS, AND NETWORKS USE COMMAND IN TERMINAL: docker system prune"
echo "2)USE THIS TO DELETE EVERYTHING: 
docker stop ascii-web
docker system prune -a --volumes"
echo "3)TO START EXISTING BUT NOT WORKING RIGHT NOW CONTAINER: docker start"
echo "4)TO STOP EXISTING AND WORKING RIGHT NOW CONTAINER: docker stop"