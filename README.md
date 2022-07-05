# message-oriented-middleware

This is a simple implementation of a calculator using ZeroMQ to implement a message oriented middleware.

Message layout:
operation val1 val2

Supported operations:
add
sub
mul
div
abs
sqrt
exp

Example of valid messages:
add 1 2     -> 3
mul -1 1    -> -1
abs -5      -> 5
sqrt 4      -> 2
exp 5 6     -> 15625

Installation:

sudo apt-get install libczmq-dev
sudo apt-get install libzmq3-dev
go get gopkg.in/zeromq/goczmq.v4
go get github.com/pebbe/zmq4

Run:
go run server/server.go
go run client/client.go
    or
go run client/client.go op val1 val2