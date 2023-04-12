FROM ubuntu:20.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    printf "1\n" | apt-get install -y cutycapt xvfb && \
    rm -rf /var/lib/apt/lists/*

ENV DISPLAY=:99

ENTRYPOINT ["sh", "-c", "Xvfb :99 -screen 0 1024x768x16 & sleep 1 && cutycapt --url=$0 --out=/app/$1"]
