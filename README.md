# Jobs and services application

Application for finding jobs and services with full cycle of service provision and final analytics by employees and their services.

## Dependencies
- [Go-Swagger](https://goswagger.io) used to generate OpenAPI specs from code & comments
    1.`docker pull quay.io/goswagger/swagger ` 
    2.`alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"`
- [Golang](https://golang.org/dl) 1.15 or higher


## Usage
### Production
- Use postgres
### Dev
- Dev ill
### Local
- (linux) `make table` create db and table in docker
- (linux) `make swagger`
- (linux) `make v3`
