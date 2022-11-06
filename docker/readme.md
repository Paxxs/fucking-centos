# 构建远程开发环境

ssh-go 是基于 ssh 的镜像并添加 golang，用于 Vscode 远程连接进入编译

```shell
docker build --rm -t local/centos:7-ssh docker/7/ssh/
docker build --rm -t local/centos:7-ssh-go-1.17 docker/7/go/
```

## 构建 docker 镜像

```shell
make docker
```
