# client-go

Go client for talking to kubernetes cluster. 

## Environment Setup
- Kind version: v0.9.0 go1.15.6 linux/amd64
- K8s Server version: v1.19.1

## Instruction for running this project
```shell
$ git clone https://github.com/pkbhowmick/client-go.git
$ cd client-go
$ go install
$ client-go version  # This command will print the version for successful installation 
```

## Pre-requisite for running Statefulset
```shell
$ client-go create-secret      # to create mongo-secret contains username & password
$ client-go create-svc         # to create headless service that maintain mongo-sts
```

## Some example commands for StatefulSet
```shell
$ client-go create-sts     # to create a stateful set "mongo-sts" of 3 replicas of mongo image
$ client-go list-sts       # to get the list of stateful set object running on default namespace
$ client-go update-sts --name=<stateful_set_name> --image=<image_name> --replicas=<number_of_replica>    # to update the statefulset object
   # Example: client-go update-sts --name=mongo-sts --replicas=4 --image=mongo
$ client-go delete-sts <sts list>     # to delete the whole stateful set
  # Example: client-go delete-sts mongo-sts  
```

