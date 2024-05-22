# 设置基础镜像
FROM golang:1.22.2-alpine AS builder

ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

# 创建工作目录
WORKDIR /app

ENV TZ Asia/Shanghai

EXPOSE 8090

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o sciproj main.go

# 创建新的构建阶段
FROM scratch as final

# 从之前的构建阶段复制文件
COPY --from=builder /app/sciproj /main
COPY --from=builder /app/config.yaml /config.yaml


CMD ["/main"]
