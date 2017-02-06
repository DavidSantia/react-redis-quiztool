# react-redis-quiztool
## Quiz Tool Demo using React, Go, Redis and Docker

To use this Quiz Tool demo, start with:
```sh
go get github.com/DavidSantia/react-redis-quiztool
go get github.com/garyburd/redigo/redis
```
As you can see above, it uses the [garyburd/redigo](https://github.com/garyburd/redigo) package, (c) Gary Burd

This demo also assumes you have [Docker](https://www.docker.com/) installed.

Next, you will need a CSV file containing your quiz data.

## Developing your own Loader
A sample plant quiz CSV is included.  To develop your own loader app:

1. Make a directory for your project
2. Use the [plant-quiz.csv](https://raw.githubusercontent.com/DavidSantia/react-redis-quiztool/master/plant-quiz.csv) as an example for how to format your quiz.
3. Create a main.go that calls New, ConnectDatastore, Parse, MapRecords, and StoreQuiz.


## How to Run the example load app
An example Load app that stores the CSV data in Redis is provided: [load/main.go](https://github.com/DavidSantia/react-redis-quiztool/blob/master/load/main.go)

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
This script builds the Go executable, and deploys it in a container with **docker-compose build**

## Launching the whole system

Finally, bring up the system:
```sh
docker-compose up
```
This launches the Redis server as a dependency to the load app.

It also mounts the [data](https://github.com/DavidSantia/react-redis-quiztool/blob/master/data) directory on the container, so that the app can access the CSV files.

