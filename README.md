### Extract files based on file feature values (md5, sha256)

[简体中文](README.zh-CN.md)

* **Hashs.xlsx**: Hash file verification value demo file:
    * md5
    * sha256

Input excel file format:

|filename|md5|sha-256|...|
|-----|-|-|-|
|a.txt|-|-|-|
|b.txt|-|-|-|
|c.txt|-|-|-|


* **Results. xlsx**: File Content Extraction Information Result Set

Output excel file format:
|fileName[verifyed]|filePath|sha-256|...|
|-----|-|-|-|
|a.txt|-|-|-|
|b.txt|-|-|-|
|c.txt|-|-|-|

### Configure&Build

**Config.ini** configuration file:

###### dirToVerify: Directory of files awaiting extract

```
#Directory of Files Awaiting Censor
dirToVerify=verify-demo
```

###### Build: build command

```
# build linux platform:
env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extraction-tool main.go
```

### Check Version

```
Usage:  -v|-V|--version: check current version
```

### Reference tool
* [excelize](https://github.com/xuri/excelize/v2)

### Important Statement

###### Prohibited for illegal purposes.
