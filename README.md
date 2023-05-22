# LineBot ChatGPT Service

This is a LINE Bot developed using Go that allows users to access the ChatGPT service. Users can engage in conversations with ChatGPT and receive intelligent responses on the LINE platform.

## Prerequisites

[Go 1.2](https://go.dev/).

[Docker](https://www.docker.com/)



## Getting Started

### 1. Register a LINE Developer Account

Before getting started, you need to register a LINE Developer account. Visit the LINE Developers website (https://developers.line.biz/) and create a new account. Under your account, create a new LINE Bot and obtain the necessary authentication credentials.

### 2. Download the Code

Download the code for this project to your local environment. You can use the Git Clone command:

```bash
git clone https://github.com/ms0529552/goLineBot.git
```

### 3. Install Dependencies

This project relies on several third-party libraries. You need to install these dependencies using `go mod` or any other appropriate tool.

You can simply used below command line order to install all the libs if you want to keep the go.mod in this repo.

```bash
go mod download

```


### 4.Set mongo on docker


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


### 5. Set Environment Variables

Before starting, set the following environment variables:

- `LINE_CHANNEL_SECRET`: The Channel Secret for your LINE Bot.
- `LINE_CHANNEL_TOKEN`: The Channel Access Token for your LINE Bot.
- `OPENAI_API_KEY`: Your API key for the ChatGPT service.

You can set these variables in a `.yaml` file. The default set will read the ./configs/config.yaml, you might need build one. Below is a example showed:

```yaml
name: "goLineBot"
mongo:
  port: 27017
  address: "mongodb://localhost:"
line:
  channel:
    secret:  `LINE_CHANNEL_SECRET`
    access_token: `LINE_CHANNEL_TOKEN`
openApi:
  go_line_bot:
    token: `OPENAI_API_KEY`

```
### 6. Set the Webhook URL

In your LINE Bot settings, configure the Webhook URL to point to your server's URL. You can use tools like ngrok (https://ngrok.com/) to expose your local server to the internet and obtain an accessible URL.




### 7. Usw ngrok to get https url
 
To make line platform interact with the application through webhooks, we need to use ngrok generate https url mapping to our app servers.

You can download ngrok [here](https://ngrok.com/)

After installaion of ngrok, we type

```bash
ngrok http 8080

```

to get url, and then copy the url  and paste to webhookurl field of your Line developers channel .

Make sure to add "/repeat" at the end of the URL

For example `https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/repeat`




## 8.Execute

To run the server you can type

```bash
go run main.go

```

## Restful Api introduction

### Get all the messags in the database

#### Request

`GET` /messages

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/messages

#### Response

    {
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
    }

### Get the messages of certain user 

#### Request

`GET` /messages?userId=\<userid\>

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/messages?UserId=U2dde8dd76c5cf5f13231d7abf82d1178

#### Response

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

### Send message back to line, and then make linebot send the message to all users who has sent message to the linebot.

#### Request

`POST` /send

    https://b558-2001-b011-381e-3046-1c14-8d3f-5276-3bc6.jp.ngrok.io/send

#### Request body

    {
        "message": "message you want to send"
    }


#### Response

    {
    "message": {
        "type": "text",
        "text": "message you want to send"
    },
    "success": "The message has been sent successfully"
    }
    



## File/directory notes:
     
      * `main.go`: main function that run this server.
      * `/controller`: main logics here.
      * `/service`: connect to db and do CRUDs.
      * `/models`: db schemas
      * `/mongo`: deal with mongodb connection.
      * others: are not important.




## Tutorial Recommandation

* [**Basic GO Lang**](https://michaelchen.tech/golang-programming/write-first-program/)
* [**Package in GO Lang**](https://calvertyang.github.io/2019/11/12/a-beginners-guide-to-packages-in-golang/)
* [**Build Web Application with Golang**](https://willh.gitbook.io/build-web-application-with-golang-zhtw/)
* [**Connection with mongoDB**](https://zhuanlan.zhihu.com/p/144308830)





### Reference

There are several golang libs be used, all of them are listed below:

* [gin](https://github.com/gin-gonic/gin)
* [viper](https://github.com/spf13/viper)
* [mongo driver](https://github.com/mongodb/mongo-go-driver)
* [cobra](https://github.com/spf13/cobra)
