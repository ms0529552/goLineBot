# goLineBot

GoLineBot is a web-based service created by golang, some go-libs, and linebot-sdk to deal with basic message sending and receiving between users and linebot. 

## Getting started

### Prerequisites

[Go](https://go.dev/)1.2.

[Docker](https://www.docker.com/)

### Installation

There are several golang libs be used, all of them are listed below:

* [gin](https://github.com/gin-gonic/gin)
* [viper](https://github.com/spf13/viper)
* [mongo\ driver](https://github.com/mongodb/mongo-go-driver)
* [cobra](https://github.com/spf13/cobra)


You can simply used below command line order to install all the libs if you want to keep the go.mod in this repo.

```bash
go mod download

```

Otherwise, using below command line order to install these go libs.

```bash
go init

go get -u github.com/gin-gonic/gin 

go get github.com/spf13/viper

go get -u github.com/spf13/cobra@latest

go get go.mongodb.org/mongo-driver/mongo

```

And for the project, we need to run mongo with docker(with mongo:4.4).

```bash
docker run -d --name <container name you want> -p <port on host>:<port on container> mongo:4.4

```

For example:

```bash
docker run -d --name mongo_container -p 27017:27017 mongo:4.4


```

Then using below order to check if container is running.

```bash
docker ps

```


### To be contiuned...