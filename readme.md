
To generate source code from .proto files, first install protoc, the protocol buffer compiler (https://grpc.io/docs/protoc-installation/). Then install two more modules to help protoc generate source code specific to the Go language:


```

apt install -y protobuf-compiler
protoc --version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

```
go get github.com/saidvandeklundert/microservices-proto/golang/order
go get github.com/sirupsen/logrus
go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
go get google.golang.org/grpc
go get google.golang.org/grpc/reflection
```

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
When in the same folder:
```
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    payment.proto
```

Get MySQL from https://www.docker.com/:
```
docker pull mysql
```

Start it using:
```
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=verysecretpass -e MYSQL_DATABASE=order mysql
```

Then:
```
CREATE DATABASE `order`;
SHOW DATABASES;
```


```
sudo apt install mysql-client
mysql -h localhost -u root -p

```
Start main using:
```
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order \
APPLICATION_PORT=3000 \
ENV=development \
go run cmd/main.go
```
https://github.com/huseyinbabal/microservices