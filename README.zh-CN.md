### Files-Extract-Tool : 根据文件特征值(md5、sha256）提取文件

* hashes.xlsx: 哈希文件校验值测试文件:
    * md5
    * sha256

Input excel file format:

|file_name|md5|sha-256|[ ... ]|
|---------|---|-------|-------|
|a.txt|-|-|-|
|b.txt|-|-|-|
|c.txt|-|-|-|


* results.xlsx: 文件内容提取信息结果集

Output excel file format:
|file_name(verifyed)|file_path|sha-256|[ ... ]|
|-----|-------------|---------|---------------|
|a.txt|verify-demo/a.txt|-|-|
|b.txt|verify-demo/test1/a.txt|-|-|
|c.txt| verify-demo/test1/test2/a.txt|-|-|

### 配置和构建

**config.ini** 配置文件：

###### dirToVerify: Directory of files awaiting extract

```
# Directory of files awaiting extract（待提取文件目录）
dirToVerify=verify-demo
```

###### build: 构建命令
```
# build linux platform:
env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extract-tool main.go
```

### 查看版本

```
Usage:  -v|-V|--version: check current version
```

### 引用鸣谢

[excelize](https://github.com/xuri/excelize/v2)

### 重要声明

###### 禁止用于非法目的！
