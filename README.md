# Getting Start with hydroponic BE

## Requirement
- Golang 1.22.1
- PostgreSQL
- MQTT emq account

## How to run
- download emqxsl-ca.crt file and place to root of this project
- run `cp .config.yaml.example .config.yaml`
- update config for database and emq account on .config.yaml file
- run command `go mod tidy` to install dependencies
- run command `go run cmd/hidroponic/main.go migration:migrate` to run migration
- run hydroponic BE using command `go run cmd/hidroponic/main.go serve`
- the default username is `admin` and the password is `Password123`
