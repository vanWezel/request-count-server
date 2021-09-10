# Request count server
This server will count the requests per instance and for the cluster in total. Built for a docker swarm cluster, because networking in swarm is very easy to configure.

## ENV variables
- PORT=8083 -> http port the server listens to
- REDIS_HOST=redis:6379 -> redis host connection
- REDIS_DB=0 -> which redis database to use
- TEST_REDIS_HOST=127.0.0.1:6379 -> for testing, redis host connection
- TEST_REDIS_DB=1 -> for testing, which redis database to use

## Dependencies
- go: 1.17
- redis: 6.2

## Packages
### [Testify](github.com/stretchr/testify)
A good package for unit tests

###[Echo](https://echo.labstack.com/guide/)
Well documented, lightweight, easy to use and unit test

###[Redigo](https://github.com/gomodule/redigo)
Well documented, lightweight
