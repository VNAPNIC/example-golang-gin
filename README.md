# Healthcare_panel
The server of the application Healthcare panel use golang

go mod

```cmd
// ini
go get gopkg.in/ini.v1

// framework
go get -u github.com/labstack/echo/v4

// Validator
go get github.com/go-playground/validator/v10

// swaggo
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/echo-swagger

swag init -g server.go
http://localhost:1323/swagger/index.html

// MySql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```
