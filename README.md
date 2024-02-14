A toy project for playing with go language and gin framework.


How to run?
---
In the gin_frontend directory, run: `npm run dev`.

In the gin_backend directory, run: `make`,

Then open a browser and enter `http://localhost:8080`.


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