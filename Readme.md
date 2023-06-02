# Filestore server

A HTTP file server with multiple available commands. This CLI acts as a client which can add, remove, edit and list files on the server.

# Getting started

- First clone the repo
- run `go build`
- use `./filestore-server`to run the server

### If you are using docker then follow below steps
- `docker build . -t filestore-server`
- `docker run -e PORT=8080 -p 8080:8080 filestore-server`

In both the cases you will be able to access the server at http://localhost:8080