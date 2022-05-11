FROM golang:bullseye

RUN useradd -ms /bin/bash admin
USER admin

RUN apt-get update && apt-get install -y \ 
    git \
    vim


