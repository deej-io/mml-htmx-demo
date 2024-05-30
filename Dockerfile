FROM golang:1.22 as gobuild

WORKDIR /api

COPY ./api/go.mod ./api/go.sum ./

RUN go mod download

COPY ./api/main.go ./
COPY ./api/assets ./assets
COPY ./api/mml ./mml
COPY ./api/components ./components

RUN go build -o /bin/main

FROM node:22 as nodebuild

WORKDIR mml-server

COPY ./mml-server/package*.json ./

RUN npm install

COPY ./mml-server/src ./src

FROM node:22

EXPOSE 8080
EXPOSE 8081

COPY --from=gobuild ./bin/main ./
COPY --from=nodebuild . ./
COPY ./entrypoint.sh .

ENTRYPOINT ./entrypoint.sh
