# THULAC分词的go语言调用
### 安装
```bash
go get -u  github.com/goylord/gothulac
```
### 调用方法
```go
import (
  " github.com/goylord/gothulac"
)

modelPath := "/program/python/python-git/thulac/models"
userDictPath := ""
preAllocSize := 1024 * 1024 * 16
t2s := 0
justSeg := 0

LoadModel(modelPath, userDictPath, preAllocSize, t2s, justSeg)
words, err := CutSentence(k)
```