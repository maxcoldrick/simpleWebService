FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o simpleWebService

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/simpleWebService /app

EXPOSE 8082
ENTRYPOINT ./simpleWebService