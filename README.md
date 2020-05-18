# gin_basic

# gin: a web framework in golang

[gin repo](https://github.com/gin-gonic/gin)

# reference documentaion

[gin 教學文章](https://ithelp.ithome.com.tw/articles/10222303)
# introduction

Gin is a web framework written in Go (Golang)

## how to start

### 1 installation

```go===
go get -u github.com/gin-gonic/gin
```

### 2 import gin into main packaage
```go===
import "github.com/gin-gonic/gin"
```

### 3 if we use some constant like http.statusOK we will import "net/http"

```go===
import "net/http"
```
### 4 test 

測試資料夾下的所有 *_test.go
```go===
go test -v ./...
```