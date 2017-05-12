Description
-----------

This is a project to create a really simple REST service using Golang.

Next, I want to explore more around the Golang concurrency, and adapt the service into something that is highly scalable e.g. using the Master/Worker pattern.

Technologies
------------

Golang, Golang libraries (net/http, encoding/json) and Gorilla Mux package for routing.

https://tour.golang.org/welcome/1

http://www.gorillatoolkit.org/pkg/mux

https://github.com/gorilla/mux

Setup
-----

1. Clone this repository
2. Set GOPATH environment variable to [path_to_cloned_repository]
3. Execute:
```
go run [path_to_cloned_repository]/src/spiritedtechie.com/user/user-http.go
```
Binary
------

To create and run a binary:
```
cd <path_to_cloned_repository>/src/spiritedtechie.com/user
go install
cd <path_to_cloned_repository>
./bin/user
```
