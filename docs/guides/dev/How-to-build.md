# How to build

Moonlight 目前之前在 MacOS（intel 和 m1都已经验证）、Linux 进行编译，暂不支持在 Windows 上的编译（不过理论上可以在 Windows WSL 中进行尝试）。

## 前置依赖

- Git (v2.8.0 or higher)
- Go (v1.17)
- rdkafka
- libgit2

## 依赖安装
在 MacOS 上安装依赖，Linux 可以参考每个依赖的文档  
rdkafka 安装 [installing-librdkafka](https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka)
```
brew install librdkafka pkg-config
```

libgit2 安装，需手动安装 1.3.0 版本
```
git clone https://github.com/libgit2/libgit2.git
git checkout v1.3.0
mkdir -p build && cd build
cmake .. -DCMAKE_INSTALL_PREFIX=/usr/local
cmake --build . --target install
```

gohub 安装

```
go get -u github.com/ping-cloudnative/moonlight-utils/tools/gohub
go install github.com/ping-cloudnative/moonlight-utils/tools/gohub
```

goimports 安装

```
go get -u golang.org/x/tools/cmd/goimports 
go install golang.org/x/tools/cmd/goimports 
```

## 编译

Moonlight 在编译前需要预生成 proto-go 和另外的部分代码， 执行

```
make proto-go-in-local
make prepare  
```

编译所有模块

> 在 m1 编译前，需要使用 dynamic 编译参数： export GO_BUILD_OPTIONS="-tags dynamic"
```
make build-all
```

编译单个模块，以 erda-server 为例

```
MODULE_PATH=erda-server make build-cross 
```
Note: Replace <module_name> with the name of the module to be build. All included module names can be found in the  [cmd](/cmd) directory.