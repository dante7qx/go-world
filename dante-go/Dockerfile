FROM centos:latest
LABEL Author="dante"

WORKDIR /app

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

## Mac 下，打包命令 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
## 参考：https://blog.csdn.net/qq_28478281/article/details/84028796
ADD ./main ./main

EXPOSE 8899

ENTRYPOINT [ "./main" ]