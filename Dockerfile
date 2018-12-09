FROM golang:1.11.2
MAINTAINER chendaole "1174250185@qq.com"

#拷贝工程到docker中
COPY . /data

#设置环境变量
ENV GOPATH /data

#设置工作目录
WORKDIR /data/src/myapp/src/

#构建工程
RUN go build main.go

#前置运行服务
CMD ["./main"]

#暴露80端口
EXPOSE 8081