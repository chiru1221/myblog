# stage 1
FROM golang:1.17.6 AS builder
COPY bap_back /go/src
WORKDIR /go/src/general
# set env for build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# build
RUN ["go", "install"]
RUN ["go", "build", "-o", "server"]

# stage 2
FROM alpine:latest
RUN ["apk", "update"]
RUN ["apk", "add", "curl"]
COPY --from=builder /go/src/general/server ./
ENV MONGO_INITDB_ROOT_USERNAME_FILE=/run/secrets/mongo-root
ENV MONGO_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongo-root-password
ENV SLACK_SS_FILE=/run/secrets/slack-signing-secret
ENV SLACK_TOKEN_FILE=/run/secrets/slack-bot-user-oauth-token
ENV GOOGLE_APPLICATION_CREDENTIALS=/run/secrets/drive-api-service-account
ENV GODEBUG=http2debug=1
EXPOSE 8080
LABEL "project"="bap"
ENTRYPOINT ["./server"]
