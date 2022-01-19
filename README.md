# Healthcare_panel
The server of the application Healthcare panel use golang

go mod

```cmd
// ini
go clean --modcache
go mod init healthcare-panel
go mod edit -module healthcare-panel
go mod tidy
go get gopkg.in/ini.v1

// framework
go get -u github.com/gin-gonic/gin

// Validator
go get github.com/go-playground/validator/v10

// swaggo
go get -u github.com/swaggo/swag/cmd/swag
swag init -g server.go

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

swag init -g server.go
http://localhost:1323/swagger/index.html

// MySql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```
