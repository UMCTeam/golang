FROM golang:1.11.2
MAINTAINER chendaole "1174250185@qq.com"

#配置环境变量
ENV GOPATH /data

#拷贝工程到docker中
COPY . /data/src/myapp

#设置工作目录
WORKDIR /data/src/myapp

#安装包管理
RUN go get -u github.com/kardianos/govendor

#安装依赖
RUN /data/bin/govendor sync

#构建工程
RUN go build src/main.go

#前置运行服务
CMD ["./main"]

#暴露80端口
EXPOSE 80