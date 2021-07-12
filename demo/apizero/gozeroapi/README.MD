## api:
1. 获取 csv 文件读取并插入数据库 post
2. 获取用户日志信息，并插入数据库 post
3. 根据需求，整合数据并显示在前端 get

## rpc
* 本地需要启动 etcd 服务
* 获取 etcd 的服务key-value ```etcdctl get /testdir/testkey --prefix```
* etcd 的 key-value 会在服务断开后一定时间能清理；etcd 被赋值的时候添加的过期时间；rpc 服务会定期的发送心跳维护键值对;
```
// example
protoc --go_out=plugins=grpc:. ./HelloWorld.proto
```
注意：
添加环境变量
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
## 如何启动 apizero 服务
1. 启动里面的 rpc 服务
2. 启动 api 服务