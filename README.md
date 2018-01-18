# macaron-swagger

[macaron](https://github.com/go-macaron/macaron) middleware to automatically generate RESTful API documentation with Swagger 2.0.

This toolkit is based on the work from https://github.com/swaggo/gin-swagger


# Usage

```
cd $project

# you need to install goswagger first
# go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger generate spec -o ./swagger.json

go run $GOPATH/src/github.com/fengbeihong/macaron-swagger/cmd/main.go init swagger.json

# now you have the file docs/docs.go
```
