FROM golang:1.15.3-alpine3.12 AS builder
WORKDIR /
COPY . .

RUN go build -o /gravity-adapter-rest ./cmd/gravity-adapter-rest/gravity-adapter-rest.go

FROM alpine:3.12.7
RUN apk update && apk upgrade --available
COPY --from=builder /gravity-adapter-rest /gravity-adapter-rest
COPY ./configs /configs
COPY ./settings /settings
COPY ./build/docker/startup.sh /startup.sh

RUN chmod 700 -R /settings && chown -R 1001  /settings
USER 1001

CMD ["sh", "/startup.sh"]
