# message-oriented-middleware

This is a simple implementation of an integer calculator using ZeroMQ to implement a message oriented middleware.

Message layout:
operation val1 val2

Supported operations:
add
sub
mul
div

Example of valid messages:
add 1 2     -> 3
add 1 2 3   -> 6
mul -1 1    -> -1

Installation:

sudo apt-get install libczmq-dev
sudo apt-get install libzmq3-dev
go get gopkg.in/zeromq/goczmq.v4
go get github.com/pebbe/zmq4

Run:
Server:
go run server/server.go

Client:
go run client/client.go op val1 val2
go run client/client.go op val1 val2 val3 ...
