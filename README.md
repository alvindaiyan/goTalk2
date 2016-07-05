# goTalk2 [![Build Status](https://travis-ci.org/alvindaiyan/goTalk2.svg?branch=master)](https://travis-ci.org/alvindaiyan/goTalk2)


This is a simple p2p chat program using [grpc](http://www.grpc.io/).

## How to Build
Run ```go build``` under the /client_app folder

### System Requirements
Go 1.6

### Dependencies
Run followings before you build to get all the dependencies.
- ```go get github.com/golang/protobuf/proto```
- ```go get golang.org/x/net/context```
- ```go get google.golang.org/grpc```


## To Start
After you install the program, you can run it with these command
- ```h``` shows the program arguments description.
- ```port ``` (default to 10000) is the port running for your local program. When another program trying to connect to you, they will need use this port on your machine. 
- ```server_addr``` (default is 127.0.0.1:10000) the target server. This specifies which target server you want to connect to. 
- ```title``` is the nick name you want to use when chat with other people.


## Architecture


	--------------------                              --------------------
	|  goTalk2 client  |------------ 				  |  goTalk2 client  |
	|				   |		  |-|-----------------|				     |
	--------------------		  |	|				  --------------------
	|				   |----------  |                 |				     |
	|	goTalk2 Server |            ----------------- |	goTalk2 Server   |
	|				   |                              |				     |
	--------------------                              --------------------
	    app 1												app2

This is a very simple idea make each program both server and client. The client' repsonsiblity is to send message and display message. The server is to receving message and pass to a client to display. 
