# goLineBot

GoLineBot is a web-based service created by golang, some go-libs, and linebot-sdk to deal with basic message sending and receiving between users and linebot. 

## Prerequisites

[Go 1.2](https://go.dev/).

[Docker](https://www.docker.com/)



## Installation

### Golang libs

There are several golang libs be used, all of them are listed below:

* [gin](https://github.com/gin-gonic/gin)
* [viper](https://github.com/spf13/viper)
* [mongo driver](https://github.com/mongodb/mongo-go-driver)
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

### Docker access method


And for the project, we need to run mongo with docker(with mongo:4.4).

You should type below command-line order to get mongo4.4 image.

```bash
docker pull mongo:4.4

```

Then you can run the container by:


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


### ngrok

To make line platform interact with the application through webhooks, we need to use ngrok generate https url mapping to our app servers.

You can download ngrok [here](https://ngrok.com/)

After installaion of ngrok, we type

```bash
ngrok http 8080

```

to get url, and then copy the url  and paste to webhookurl field of your Line developers channel .

Make sure to add "/repeat" at the end of the URL

For example `https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/repeat`




## Execute

To run the server you can type

```bash
go run main.go

```

## Restful Api introduction

### Get all the messags in the database

#### Request

`GET` /messages

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/repeat/messages

#### Response

`{
    "messages list": [
        {
            "id": "17807235074942",
            "user_id": "U604fef6644e218cc6a1a8925391d30fe",
            "type": "",
            "text": "test8:53",
            "created_at": "2023-03-15T12:54:01.01Z"
        },
        {
            "id": "17807243183240",
            "user_id": "Uaf4f4b8ceec9e6a90d366fa93156d580",
            "type": "",
            "text": "teSt2055",
            "created_at": "2023-03-15T12:55:31.29Z"
        },
        {
            "id": "17807249427867",
            "user_id": "Uaf4f4b8ceec9e6a90d366fa93156d580",
            "type": "",
            "text": "hello",
            "created_at": "2023-03-15T12:56:40.493Z"
        },
      
    ]
}`

#### Request

`GET` /messages?userId=<userid>

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/repeat/messages?UserId=U2dde8dd76c5cf5f13231d7abf82d1178

#### Response

`
{
    "messages list": [
        {
            "id": "17807548717269",
            "user_id": "U2dde8dd76c5cf5f13231d7abf82d1178",
            "type": "",
            "text": "test1",
            "created_at": "2023-03-15T13:53:34.962Z"
        },
        {
            "id": "17807647591894",
            "user_id": "U2dde8dd76c5cf5f13231d7abf82d1178",
            "type": "",
            "text": "test2",
            "created_at": "2023-03-15T14:13:49.111Z"
        }
    ]
}
`

### Send message back to line, and then make linebot send the message to all users who has sent message to the linebot.

#### Request

`POST` /send

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/repeat/send

#### Request body
`
{
    "message": "message you want to send"
}
`

#### Response
`
{
    "message": "The message has been sent successfully"
}
`

## File/directory notes:
     
      * `main.go`: main function that run this server.
      * `/controller`: main logics here.
      * `/service`: connect to db and do CRUDs.
      * `/models`: db schemas
      * `/mongo`: deal with mongodb connection.
      * others: are not important.


### Execute:




### To be contiuned...