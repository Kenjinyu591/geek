FROM golang:1.17-alpine3.14 AS build
LABEL authors="kenfish591"

# 设置工作目录
WORKDIR /app

# 将当前目录的内容复制到工作目录
COPY . .

RUN go build -o httpsvr httpsvr.go

# 更改为使用轻量级的 alpine 基础镜像
FROM alpine:3.14

# 将编译好的程序复制到新的镜像中
COPY --from=build /app/httpsvr /app/httpsvr

RUN chmod +x /app/httpsvr

CMD ["/app/httpsvr"]