FROM node:20.2

WORKDIR /many-panel
COPY . .
WORKDIR /many-panel/frontend
ENV NODE_OPTIONS=--max_old_space_size=3000
RUN npm install && npm run build:pro


FROM golang:alpine as builder

WORKDIR /many-panel
copy --from=0 /many-panel /many-panel
WORKDIR /many-panel/cmd/server

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .


FROM alpine:latest

WORKDIR /opt/many-panel
COPY --from=1 /many-panel/cmd/server/server ./

EXPOSE 9999
ENTRYPOINT ./server
