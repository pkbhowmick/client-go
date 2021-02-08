# client-go

Go client for talking to kubernetes cluster. 

## Instruction for running this project
```shell
$ git clone https://github.com/pkbhowmick/client-go.git
$ cd client-go
$ go mod download
$ go build .
$ ./client-go version  # This command will print the version for successful installation 
```

## Some example commands for StatefulSet
```shell
$ ./client-go create-sts     # to run a stateful set of 3 replicas of mongo image 
$ ./client-go list-sts       # to get the list of stateful set object running on default namespace
$ ./client-go update-sts --name=<stateful_set_name> --image=<image_name> --replicas=<number_of_replica>    # to update the statefulset object
$ ./client-go delete-sts --name=<stateful_set_name>     # to delete the whole stateful set    
```

