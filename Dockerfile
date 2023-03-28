# 第一阶段构建编译环境
FROM golang:alpine AS builder

WORKDIR /project/travel_planning_assistant
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go generate && go env && go build -o server .
# 第二阶段构建运行环境
FROM alpine:latest

WORKDIR /project/travel_planning_assistant

COPY --from=builder /project/travel_planning_assistant ./

EXPOSE 8888

ENTRYPOINT ./server -c config.dev.yaml