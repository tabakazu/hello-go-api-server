# syntax=docker/dockerfile:1

FROM golang:1.23-bullseye AS dev

RUN apt-get update -qq && apt-get install -y build-essential jq postgresql-client unzip

CMD ["sleep", "infinity"]
