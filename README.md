# react-redis-quiztool
## Quiz Tool Demo using React, Go, Redis and Docker

To use this Quiz Tool demo, start with:
```sh
go get github.com/DavidSantia/react-redis-quiztool
```
This downloads the quiztool package, along with its dependencies:

* [github.com/garyburd/redigo](https://github.com/garyburd/redigo) package, (c) Gary Burd
* [github.com/gorilla/websocket](https://github.com/gorilla/websocket) package, (c) Gorillatoolkit.org

This demo also assumes you have [Docker](https://www.docker.com/) installed, and uses

* Docker hub [Redis](https://hub.docker.com/_/redis) container, (c) Redis.io

## Architecture

I wanted to create a simple React/JS app that presents a quiz.  I also wanted to serve the data from Redis, and keep things simple and fast by connecting Redis directly to the browser.

When I first researched this, I found various NodeJS modules that interface with Redis; however, these are all server side implementations. Since Redis uses a TCP socket, browsers don't let you interface directly. Security concerns have kept TCP sockets out of browsers. (Although Chrome does have a socket library, it keeps things secure by only allowing sockets in packaged apps, where restrictions can be specified in their manifest.)

So I created a light-weight adaptor ([redis-ws/main.go](https://github.com/DavidSantia/react-redis-quiztool/blob/master/redis-ws/main.go)) to copy onto the Redis container, providing a WebSocket interface. I also created a separate load app on a second container, that loads quiz data and then exits.  The overall architecture is as follows:
![Figure 1: Architecture](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/README-Architecture.png)

## How to Run the example load app
An example Load app ([load/main.go](https://github.com/DavidSantia/react-redis-quiztool/blob/master/load/main.go)) that stores the CSV data in Redis is provided.

### Running locally
To run this app, first you need to launch a Redis container.  At this point the standard Redis container will do.  The following maps the port Redis uses to localhost:
```sh
docker run --rm --name redis -p 6379:6379 redis:alpine
```
This will run on the terminal (until you type Ctrl-C)

On another terminal, run the load app as follows:
```sh
go run load/main.go
```

### Deploying the load app in a container
Next, we will run the whole system in Docker. To do this, first stop the currently running Redis by typing Ctrl-C on its terminal.  Then, build the load app:
```sh
./build.sh
```
This script builds two Go executables, the load app and the adaptor.  It then deploys these executables onto docker images via **docker-compose build**.

## Launching the whole system

Finally, bring up the system:
```sh
docker-compose up
```
This launches

1. The Redis server (with Websocket adaptor) container, as a dependency to the load app
2. The Load app container

It also mounts the [data](https://github.com/DavidSantia/react-redis-quiztool/blob/master/data) directory on the Load container, so that the app can access the CSV files.

## Running the React app

The React app uses a Redis container with the WebSocket adaptor attached.  This is built and launched in as explained in the section above.

Once launched, use a separate terminal, go to the [quiz](https://github.com/DavidSantia/react-redis-quiztool/blob/master/quiz) directory and follow the README.

## Developing your own Loader
A sample plant quiz CSV is included.  To develop your own loader app, start with a CSV file containing your quiz data.

1. Make a directory for your project
2. Use the [plant-quiz.csv](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/plant-quiz.csv) as an example for how to format your quiz.
3. Create a main.go that calls New, ConnectRedis, Parse, MapRecords, and StoreQuiz.

## Troubleshooting the Redis Websocket

Launching the whole system loads the sample data into Redis. If you want to test out the Redis Websocket by itself, naviage to a test page such as [websocket.org -> Demos -> Echo Test](http://websocket.org/echo.html).

From here, connect to Redis as follows:

* Make sure the "Use secure WebSocket (TLS)" box is unchecked
* Enter the address "ws://localhost:4000"
* Press Connect

You should see "CONNECTED" in the Log, as shown.
![Figure 2: Debugging Websocket](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/README-DebugWS.png)

Replace the default Message. Use a JSON struct with **command** and **data** fields as follows:

* The field "command" should contain a Redis command, as found on [redis.io/commands](https://redis.io/commands)
* The field "data" is a string containing the arguments to the command

### Examples:

Get the meta-data for Quiz 1
```json
{"command":"HGETALL", "data":"quiz:1"}
```

Get the number of questions in Quiz 1, Category 2
```json
{"command":"HGET", "data":"quiz:1:c:2 questions"}
```

## Contributors

* [DavidSantia](https://github.com/DavidSantia)
