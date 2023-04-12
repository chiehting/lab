#!/bin/bash

docker build -t cutycapt .
docker run --rm -v $(pwd):/app cutycapt https://google.com google.com.png
docker run --rm -v $(pwd):/app cutycapt https://www.youtube.com youtube.com.png
docker run --rm -v $(pwd):/app cutycapt https://medium.com medium.com.png
docker run --rm -v $(pwd):/app cutycapt https://github.com github.com.png
