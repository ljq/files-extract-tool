### Files Extract Tool: extract files based on file feature values (md5, sha256)

[简体中文](README.zh-CN.md)

* **Hashs.xlsx**: Hash file verification value demo file:
    * md5
    * sha256

Input excel file format:

|file_name|md5|sha-256|[ ... ]|
|---------|---|-------|-------|
|a.txt|-|-|-|
|b.txt|-|-|-|
|c.txt|-|-|-|


* **Results. xlsx**: File Content Extraction Information Result Set

Output excel file format:
|file_name[verifyed]|file_path|sha-256|[ ... ]|
|-----|-------------|---------|---------------|
|a.txt|verify-demo/a.txt|-|-|
|b.txt|verify-demo/test1/a.txt|-|-|
|c.txt| verify-demo/test1/test2/a.txt|-|-|

### Configure and build

**Config.ini** configuration file:

###### dirToVerify: Directory of files awaiting extract

```
#Directory of Files Awaiting Extract
dirToVerify=verify-demo
```

###### Build: build command

```
# build linux platform:
env GOOS=linux GOARCH=amd64 go build -o ./bin/files-extract-tool main.go
```

### Check Version

```
Usage:  -v|-V|--version: check current version
```

### Reference tool
* [excelize](https://github.com/xuri/excelize/v2)

### Important Statement

###### Prohibited for illegal purposes.
