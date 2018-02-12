# 17.3 libcontainerを使ったコンテナの自作
Linux上でしか動作しないので注意する。

## ブートのための下準備
```bash
$ docker pull alpine
$ docker run --name alpine alpine
$ docker export alpine > alpine.tar
$ docker rm alpine
$ mkdir rootfs
$ tar -C rootfs -xvf alpine.tar
```
