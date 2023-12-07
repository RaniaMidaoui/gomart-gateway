# gomart-gateway
This repository contains the code and Dockerfile for the API Gateway of the **goMart** application, along with the Jenkinsfile describing the CI/CD pipeline.

To run the code, you need to have Golang package installed:

1- Download the package from [the official website](https://go.dev/doc/install)
##### For Linux users:
2- Remove any previous Go installation  then extract the archive you just downloaded into /usr/local:
```
 $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```
3- Add /usr/local/go/bin to the PATH environment variable:
```
$ export PATH=$PATH:/usr/local/go/bin
```
4- Verify that you've installed Go:
```
$ go version
```
##### For Windows users:
Follow the prompt after opening the MSI file you downloaded from [the official website](https://go.dev/doc/install).

### To run locally:
Start by making sure all the dependencies are installed and run the code, it will tell you that it's listenning on the application port configured:
```
$ go mod tidy
$ make proto
$ make server
```

