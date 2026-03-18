FROM golang:1.19.5 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev
RUN GOOS=linux go build -gcflags="all=-N -l" -o /main

FROM debian:buster
EXPOSE 8080 8081 40000
WORKDIR /
COPY --from=build-env /main .
COPY --from=build-env /dockerdev/.env .
ENTRYPOINT ./main
