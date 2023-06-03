# Filestore server

A HTTP file server with multiple available commands. This CLI acts as a client which can add, remove, edit and list files on the server.

# Getting started

- First clone the repo
- run `go build`
- use `./filestore-server`to run the server

### If you are using docker then follow below steps

- `docker build . -t filestore-server`
- `docker run -e PORT=8080 -p 8080:8080 filestore-server`

### The above commands will first take the local dockerfile and create an image and then run that image on your docker instance.

---

### Alternatively you can use the pre-built docker image and run your container from that image. To do that simply run

`docker run -e PORT=8080 -p 8080:8080 ayushsatyam146/filestore-server`

### This will pull image from docker hub and run the container

### In all the cases you will be able to access the server at http://localhost:8080

---

### Kubernetes manifest file are also present here. You can easily run the container image on a local or managed k8s cluster. Just run the following commands in a suitable environment

- `kubectl create -f deployment.yaml`
- `kubectl create -f services.yaml`
- We're using default namespace
- Run `kubectl get svc` to check if `filestore-server` service is running and note the NodePort
- Run `minikube ip` to get the IP
- You can access the server at `minikube-ip:nodePort `
- Run `kubectl get pods` to get pod description
