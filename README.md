How to upgrade?
---
To upgrade `gin`, run:
```
go get -u github.com/gin-gonic/gin
```

To upgrade all dependencies, run:
```
go get -u all
```


Downloading go packages quickly
---
Run the following commands:
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

Next step
---
See https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/