## HTTP Service
Check [api.yaml](api/api.yaml) for the API documentation or upload it to https://editor.swagger.io/

### Features:
* Implementation of Clean Architecture in Go
* Dependency injection
* Context awareness
* Structure of the Project follows the [Project Layout](https://github.com/golang-standards/project-layout).

### Motivation:
* API first 
* Keep things simple
* Codegen tools and frameworks to speed up the development (I had not much of a free time)

### Structure:
* `api/*` - Openapi documentation and codegen tools and generated code
* `cmd/*` - main application
* `internal/*` - business-logic

### 3rd party libraries
* Echo HTTP framework
* Viper for configuration
* Cobra CLI
* [JWT](https://github.com/golang-jwt/jwt/v5) library
* [Abstract JSON](https://github.com/spyzhov/ajson) to analyse dynamic JSON structures
* Mockery 

### Test
* `go test ./...`


### Contribute
* Check prerequisites
* Add changes to [api.yaml](api/api.yaml)
* run `go generate`

### Development Prerequisites
* Install [mockery](https://github.com/vektra/mockery)
* Install [oapi-codegen](https://github.com/deepmap/oapi-codegen)

### Build & Run
* `go build -o app ./cmd`
* `./app`
* OR `go run cmd/main.go`

### Docker
* `docker build -t maksym-code-assignment .`
* `docker run -p8280:8280 maksym-code-assignment`


### Usage
Service should be listening ` [::]:8280`

Get Auth token 
```shell
curl --location 'http://localhost:8280/auth' \
--header 'Content-Type: application/json' \
--data '{
    "username": "maksym",
    "password": "trofimenko"
}'
```
The result:
`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtYWtzeW0iLCJleHAiOjE2ODA1MTQ3NDB9.GS6bPVYHqDP8XgrqAVl8_yFD1frBcLLOnQA5x2Tu3uY`
* Use https://jwt.io/ to check if JWT has a subject `maksym`
* Check that expiration time is in one hour
* Copy the token for a next request

Get checksum of a sum of all numbers of a

```shell
curl --location 'http://localhost:8280/sum' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtYWtzeW0iLCJleHAiOjE2ODA1MTQ3NDB9.GS6bPVYHqDP8XgrqAVl8_yFD1frBcLLOnQA5x2Tu3uY' \
--header 'Content-Type: application/json' \
--data '{"a":{"b":4},"c":-2, "d": {"a":{"b":4},"c":-2}}'
```

Result:
`4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a`

* Go to website https://codebeautify.org/sha256-hash-generator
* Enter number `4` as an input
* Check if hash sum matches.

That's it.
