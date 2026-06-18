#!/bin/bash

read -r -p "enter the name of the image " imagename && docker build -t $imagename .
read -r -p "enter name of the container " container
docker run -d -p 8080:8080 --name $container $imagename
echo "docker container $container running at port 8080 exposed from 8080"
