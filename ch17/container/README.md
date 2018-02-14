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

## Result

```bash
$ docker build -t build-container . --no-cache=true
$ docker run --privileged -it build-container /bin/ash
/go/src/github.com/budougumi0617/gsp/ch17/container # ./main
/bin/sh: can't access tty; job control turned off
/ # /bin/hostname
testing
/ # exit
/go/src/github.com/budougumi0617/gsp/ch17/container # exit
```
