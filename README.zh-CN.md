### 根据文件特征值(md5、sha256）提取文件

* hashes.xlsx: 哈希文件校验值测试文件:
    * md5
    * sha256
* results.xlsx: 文件内容提取信息结果集

### 配置 & 构建

**config.ini** 配置文件：

###### dirToVerify: Directory of files awaiting extract

```
# Directory of files awaiting censor（待提取文件目录）
dirToVerify=verify-demo
```

###### build: 构建命令
```
# build linux platform:
env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extraction-tool main.go
```

### 查看版本

```
Usage:  -v|-V|--version: check current version
```

### 引用鸣谢

[excelize](https://github.com/xuri/excelize/v2)

### 重要声明

###### 禁止用于非法目的！
