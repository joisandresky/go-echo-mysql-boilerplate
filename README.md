# Golang Echo MYSQL Boilerplate

This is My Template/Boilerplate to Create REST API/Microservice Using:

 - Go
 - Echo Framework (v4)
 - MySQL
 - Viper
 - Paseto Token
 - Migrate ([golang-migrate](https://github.com/golang-migrate/migrate))

# How to Run ?

 1. Clone this repo
 2. `go mod init YOUR_GO_MOD_NAME`
 3. `go mod download`
 3. copy `configs/env_example.yaml` and rename into `env.yaml`
 4. `go run cmd/main.go`