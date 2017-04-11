# react-redis-quiztool
## Quiz Tool Demo using React, Go, Redis and Docker

To use this Quiz Tool demo, start with:
```sh
go get github.com/DavidSantia/react-redis-quiztool
go get github.com/garyburd/redigo/redis
go get github.com/gorilla/websocket
```
As you can see above, it uses

* [garyburd/redigo](https://github.com/garyburd/redigo) package, (c) Gary Burd
* [gorilla/websocket](https://github.com/gorilla/websocket) package, (c) Gorillatoolkit.org

This demo also assumes you have [Docker](https://www.docker.com/) installed, and uses

* Docker hub [Redis](https://hub.docker.com/_/redis) container, (c) Redis.io

## Architecture

I wanted to create a simple React/JS app that presents a quiz.  I also wanted to serve the data from Redis, and keep things simple and fast by connecting Redis directly to the browser.

When I first researched this, I found various NodeJS modules that interface with Redis; however, these are all server side implementations. Since Redis uses a TCP socket, browsers don't let you interface directly. Security concerns have kept TCP sockets out of browsers. (Although Chrome does have a socket library, it keeps things secure by only allowing sockets in packaged apps, where restrictions can be specified in their manifest.)

So I created a light-weight adaptor ([redis-ws/main.go](https://github.com/DavidSantia/react-redis-quiztool/blob/master/redis-ws/main.go)) to copy onto the Redis container, providing a websocket interface. I also created a separate load app on a second container, that loads quiz data and then exits.  The overall architecture is as follows:
![Figure 1: Architecture](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/README-Architecture.png)

## How to Run the example load app
An example Load app ([load/main.go](https://github.com/DavidSantia/react-redis-quiztool/blob/master/load/main.go)) that stores the CSV data in Redis is provided.

### Running locally
To run this app, first you need to launch a Redis container.  The following maps the port Redis uses to localhost:
```sh
docker run -it --rm --name redis -p 6379:6379 redis:alpine
```
This will run on the terminal (until you type Ctrl-C)

On another terminal, run the load app as follows:
```sh
go run load/main.go
```

### Deploying the app in a container
Next we will run the whole system in Docker. To do this, first stop the currently running Redis by typing Ctrl-C on its terminal.  Then, build the load app:
```sh
./build.sh
```
This script builds the two Go executables, and deploys them onto their containers with **docker-compose build**

## Launching the whole system

Finally, bring up the system:
```sh
docker-compose up
```
This launches

1. The Redis server (with websocket adaptor) container, as a dependency to the load app
2. The Load app container

It also mounts the [data](https://github.com/DavidSantia/react-redis-quiztool/blob/master/data) directory on the Load container, so that the app can access the CSV files.

## Developing your own Loader
A sample plant quiz CSV is included.  To develop your own loader app, start with a CSV file containing your quiz data.

1. Make a directory for your project
2. Use the [plant-quiz.csv](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/plant-quiz.csv) as an example for how to format your quiz.
3. Create a main.go that calls New, ConnectRedis, Parse, MapRecords, and StoreQuiz.

