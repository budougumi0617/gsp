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
$ docker build -t build-container .
$ docker run -it build-container /bin/ash
/go/src/github.com/budougumi0617/gsp/ch17/container # sudo ./main
WARN[0000] signal: killed
2018/02/14 00:43:37 container_linux.go:348: starting container process caused "process_linux.go:279: applying cgroup configuration for process caused \"mkdir /sys/fs/cgroup/cpuset/system: read-only file system\""
```
