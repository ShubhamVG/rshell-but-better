# rshell-but-better

This is my first _big_ project I made while learning Go. It started off as an rshell but soon was transferred into something like a remote access tool that can connect to multiple clients and execute terminal commands on it. The client can connect to the server via TCP and communicate. The server has multiple goroutines and each goroutine does its own job without relying on each other **(as of now assuming I am missing something)**. The address and port of where to connect is configured when the script is built into an executable so look into the `backend/cmd` directory to know more about it.

**NOTE:** There is no encryption whatsoever. This project was made for fun when I was learning about the great Go language so it was made WITHOUT the intent of ever being used. I may add encryption in the future when I will be learning number theory and the mathematics behind encryption but it will still be without the intention of using this project.

## Backend

The backend is entirely written in Go. As present in the `backend/cmd` directory, there are two main components. A server that has connect to multiple clients and send them commands and receive outputs (of stdout) from them and such. There is also a client which basically runs the commands or follows any other sort of instruction given and sends them back.

### Server

The server has multiple goroutines. One to accept connection, one to send data and one to receive data. And unless I am forgetting something, that's it.
The responses that it received from the connected client get stored in a bucket of capacity 10 by default, i.e., it can only store the latest 10 responses received.

### Client

The client works on one goroutine only. Just to receive data which are terminal commands, then parse the command and do whatever it says and send back any output (from stdout).

## Frontend

~~There is also a frontend part which is written* (actually being written as of now) using Flutter. It is technically just a client but written in dart and its purpose is to act as an admin monitor as it would be awful to run a GUI from the shared server like a tiny VPS.~~

Update: Due to limitations and moving priorities, there will be no progress on the frontend part and I am removing it from the repo as it is incomplete. The biggest limitation was the TCP communication between two different programming languages. They don't use buffers so the frontend (written in dart) was receiving multiple pieces of data at the same time and this could have been easily solved by adding a "limiter" character or something but as I am working on a different project right now, I will not be working on the frontend of this and hence it is being removed.

## TODO(s)
- Complete the frontend.
- Maybe add encryption.

Thanks for reading till here. Drop me a hi on Discord or email me. (And yes, I do read my emails everyday.)