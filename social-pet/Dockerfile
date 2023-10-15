FROM golang:latest

# 维护者
MAINTAINER ""

#创建工程文件夹
RUN mkdir -p /app
RUN mkdir -p /app/conf
RUN mkdir -p /app/logs
RUN mkdir -p /app/static


# 拷贝当前目录代码到镜像
COPY ./conf /app/conf
COPY ./static /app/static
COPY ./social-pet /app/social-pet

#设置时区
COPY ./Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /app

VOLUME ["/app/conf","/app/logs", "/app/static"]

ENTRYPOINT ["./social-pet"]
