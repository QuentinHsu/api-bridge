# 使用最新的 golang 镜像作为基础镜像
FROM debian:latest
# 在容器中创建一个目录来存放我们的应用
RUN mkdir /app
# 将工作目录切换到 /app 下
WORKDIR /app
# 将本地当前目录的所有文件拷贝到 /app 目录
COPY . /app
# 在容器启动时运行命令
CMD ["/app/main"]
